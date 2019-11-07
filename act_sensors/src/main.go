package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func parseTime(text string) int64{
	parsedStime, _ := time.Parse(
		"2006-01-02 15:04:05",
		text)

	return parsedStime.Unix()
}

func openActivities(ch *amqp.Channel, q amqp.Queue){
	file, err := os.Open("activity.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	time.Sleep(2500)
	for scanner.Scan() {
		body := scanner.Text()
		//fmt.Println(body)
		activityReg := regexp.MustCompile(`(.+?)\t+(.+?)\t+(.+?)\t+`)
		matches := activityReg.FindStringSubmatch(body)

		jsonEntry := fmt.Sprintf(`{"patient_id":1, "activity":"%s", "start": %d, "end": %d}`, matches[3], parseTime(matches[1]), parseTime(matches[2])) //Build connection string

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(jsonEntry),
			})
		log.Printf(" [x] Sent %s", jsonEntry)
		failOnError(err, "Failed to publish a message")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
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

	openActivities(ch, q)
}
