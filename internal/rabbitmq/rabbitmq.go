package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Connect() error
}

type RabbitMQ struct {
	Conn *amqp091.Connection
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

	return nil
}
