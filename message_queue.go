package main

import (
	"github.com/streadway/amqp"
)

const (
	cab_queue_name = "cab_queue"
)

func InitMessageQueue() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	queue, err := ch.QueueDeclare(
		cab_queue_name, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		queue.Name, // queue
		"", // consumer
		true, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for message := range messages {
			MessageReceived(message)
		}
	}()
}

// This function is only for one time send for now
func SendMessage(message []byte) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cab_queue_name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        message,
		})
	failOnError(err, "Failed to publish a message")
}

func MessageReceived(message amqp.Delivery) {
	LogInfo("Received a message: %s", message.Body)
	NewDbQuery(IndexName).Put(NewCabFromJson(message.Body))
}

