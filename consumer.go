package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RecivedFromRabbitmq(queue string, conn *amqp.Connection) error {
	channel, err := conn.Channel()
	if err != nil {
		log.Println("failed to create channel", err)
	}

	defer conn.Close()

	defer channel.Close()

	q, err := channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Println("failed to declare a queue", err)
		return nil
	}

	// We use exchange when we want producer to send to different queues without interacting directly with queue
	err = channel.ExchangeDeclare(
		"account_exchange", // Exchange name
		"fanout",           // Exchange type
		true,               // Durable
		false,              // Auto-deleted
		false,              // Internal
		false,              // No-wait
		nil,                // Arguments
	)
	if err != nil {
		log.Println("Exchange declaration failed", err)
		return nil
	}

	// Bind the queue to the exchange to let them know each other. with that we can have as many queues as we want to the same exchange
	err = channel.QueueBind(
		q.Name,             // Queue name
		"",                 // Routing key
		"account_exchange", // Exchange
		false,
		nil,
	)
	if err != nil {
		log.Println("queue bind failed", err)
	}

	msgs, err := channel.Consume(
		q.Name, // routing key
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // mandatory
		false,  // immediate
		nil,
	)

	if err != nil {
		log.Println("failed to consume", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Recived a message: %s", d.Body)
		}
	}()

	<-forever

	return nil
}
