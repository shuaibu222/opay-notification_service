package main

import (
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// pulsar connection
	rabbitClient, err := connect()
	if err != nil {
		log.Println("Error connecting to pulsar from consumer: ", err)
	}

	defer rabbitClient.Close()

	// Account related consumer

	err = RecivedReviewToRabbitmq("accountCreated", rabbitClient)
	if err != nil {
		log.Println("accountCreated error: ", err)
	}
	err = RecivedReviewToRabbitmq("accountDeleted", rabbitClient)
	if err != nil {
		log.Println("accountDeleted error: ", err)
	}
	err = RecivedReviewToRabbitmq("accountUpdated", rabbitClient)
	if err != nil {
		log.Println("accountUpdated error: ", err)
	}

	// Transaction related consumer

	err = RecivedReviewToRabbitmq("transaction", rabbitClient)
	if err != nil {
		log.Println("Error consuming from consumer: ", err)
	}

}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue until rabbit is ready
	for {
		client, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = client
			break
		}

		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
