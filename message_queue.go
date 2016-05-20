package main

import (
	"github.com/streadway/amqp"
)

type AMQPMessager struct {
	url string
	connection *amqp.Connection
	channel *amqp.Channel
}

type AMQPSender struct {
	messager *AMQPMessager
	queue amqp.Queue //check this pointer out
}

func NewAMQPMessager(url string) *AMQPMessager {
	messager := &AMQPMessager{
		url: url,
	}
	return messager
}

func (messager *AMQPMessager) NewAMQPSender(queue_name string) *AMQPSender {
	queue, err := messager.channel.QueueDeclare(
		queue_name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	sender := &AMQPSender{
		messager: messager,
		queue: queue,
	}

	return sender
}

func (sender *AMQPSender) Send(message []byte) *AMQPSender {
	err := sender.messager.channel.Publish(
		"",     // exchange
		sender.queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        message,
		})
	failOnError(err, "Failed to publish a message")

	return sender
}

func (messager *AMQPMessager) Connect() *AMQPMessager {
	connection, err := amqp.Dial(messager.url)
	failOnError(err, "Failed to connect to RabbitMQ")
	messager.connection = connection

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	messager.channel = channel

	return messager
}

func (messager *AMQPMessager) Disconnect() *AMQPMessager {
	messager.channel.Close()
	messager.connection.Close()
	return messager
}


type consumeCallback func(amqp.Delivery)
func (messager *AMQPMessager) Consume(queue_name string, callback consumeCallback) *AMQPMessager {
	queue, err := messager.channel.QueueDeclare(
		queue_name, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := messager.channel.Consume(
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
			callback(message)
		}
	}()

	return messager
}
func MessagerCabUpdate(message amqp.Delivery) {
	LogInfo("Received a message: %s", message.Body)
	NewDbQuery(IndexName).Put(NewCabFromJson(message.Body))
}
func InitMessageQueue() {
	amqpMessager = NewAMQPMessager(messager_queue_url).Connect()
	amqpMessager.Consume(messager_cab_queue_name, MessagerCabUpdate)
	amqpSender = amqpMessager.NewAMQPSender(messager_cab_queue_name)
}

// This function is only for one time send for now
func SendMessage(message []byte) {
	amqpSender.Send(message)
}

