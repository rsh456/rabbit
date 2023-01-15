package rabbitmq

import (
	"fmt"

	amqp "github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (rmq *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	rmq.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	fmt.Println("Successfully Connected to RabbitMQ")
	rmq.Channel, err = rmq.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = rmq.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	return nil
}

// Publish - takes in a string message and publishes to a queue
func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully published a message to the queue")
	return nil
}

// coonsume - consumes messages from our test queues
func (r *RabbitMQ) Consume() {
	fmt.Println("Cosuming")
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Received Messages: %s\n", msg.Body)
	}
}

// New RabbitMQService - rturns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}
