package controllers

import (
	"dslic/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ListenForRabbits() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//msg := fmt.Sprintf("%s", d.Body)
			activity := &models.Activity{}
			_ = json.Unmarshal(d.Body, activity)
			//_ := json.NewDecoder(msg).Decode(activity)
			log.Printf("Received a message: %s", d.Body)

			message := ""
			if activity.Name == "Sleeping" && activity.End-activity.Start > 72000 {
				message = fmt.Sprintf("Patient #%d has slept too much!!", activity.ID)
			}
			if activity.Name == "Outdoor" && activity.End-activity.Start > 72000 {
				message = fmt.Sprintf("Patient #%d has been outdoor too much!!", activity.ID)
			}
			if activity.Name == "Toileting" && activity.End-activity.Start > 60000 {
				message = fmt.Sprintf("Patient #%d has stomach problems!!", activity.ID)
			}

			if message != "" {
				log.Print("Naughty patient!!!")
				for _, conn := range wsConnections {
					_ = conn.WriteMessage(websocket.TextMessage, []byte(message))
				}
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
