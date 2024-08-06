package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Saqlainabbasi/api/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService interface {
	Pubblisher(ctx context.Context, eventName string, eventMessage []byte) error
}

type rabbitmq struct {
	conn *amqp.Connection
}

var uri = utils.EnvAMQPKEY()

// Constructor function.....
func NewAMQP() RabbitMQService {

	connection, err := amqp.Dial(uri)
	if err != nil {
		defer log.Fatalf("RABBITMQ NOT CONNECTED: %s", err.Error())
		defer connection.Close()
		return nil
	}

	defer log.Println("RABBITMQ CONNECTED")
	return &rabbitmq{conn: connection}
}

func (mb *rabbitmq) Pubblisher(ctx context.Context, evenntName string, eventMessage []byte) error {
	channel, err := mb.conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		evenntName, // name
		false,      // durable
		false,      // auto delete
		false,      // exclusive
		false,      // no wait
		nil,        // args
	)
	if err != nil {
		panic(err)
	}

	// publishing a message
	err = channel.Publish(
		"",         // exchange
		evenntName, // key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte("Test Message"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
	return nil
}
