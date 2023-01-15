package main

import (
	"fmt"
	rabbitmq "rabbit/internal"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go Rabbit MQ")

	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("Hi")
	if err != nil {
		return err
	}

	app.Rmq.Consume()

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up the application")
		fmt.Println(err)
	}
}
