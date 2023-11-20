package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail() {
	// sending email
	m := gomail.NewMessage()
	m.SetHeader("From", "shuaibuabdulkadir222@gmail.com")
	m.SetHeader("To", "shuaibuabdulkadir656@gmail.com")
	m.SetHeader("Subject", "Assalamu alaikum")
	m.SetBody("text/html", "Hi <b>Shuayb</b>!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "shuaibuabdulkadir222@gmail.com", "fjnmxuxrfvfrunrd")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email", err)
	}
}
