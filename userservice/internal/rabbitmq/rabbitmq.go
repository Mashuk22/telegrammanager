package rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewService() *RabbitMQ {
	return &RabbitMQ{}
}
func (rmq *RabbitMQ) Connect() error {
	var err error
	rmq.Conn, err = amqp091.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		return err
	}
	rmq.Channel, err = rmq.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = rmq.Channel.QueueDeclare("UserService", false, false, false, false, nil)
	if err != nil {
		return err
	}

	fmt.Print("Successfully published message to the queue")
	return nil
}

func (rmq *RabbitMQ) Publish(message string) error {
	err := rmq.Channel.Publish("", "UserService", false, false, amqp091.Publishing{ContentType: "text/plain", Body: []byte(message)})
	if err != nil {
		return err
	}
	return nil
}

func (rmq *RabbitMQ) Consume() {
	messages, err := rmq.Channel.Consume("UserService", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	for message := range messages {
		fmt.Printf("Received message: %s\n", message.Body)
	}
}
