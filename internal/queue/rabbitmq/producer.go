package rabbitmq

import (
	"demo-smtp/internal/config"
	"fmt"
	"log/slog"

	"github.com/streadway/amqp"
)

type Producer[T any] struct {
	QueueName string
	Channel   *amqp.Channel
}

func NewProducer[T any]() *Producer[T] {
	connection, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%d/",
			config.MainConfig.RabbitMQ.Username,
			config.MainConfig.RabbitMQ.Password,
			config.MainConfig.RabbitMQ.Host,
			config.MainConfig.RabbitMQ.Port,
		),
	)

	if err != nil {
		slog.Error("Failed to connect to RabbitMQ: " + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		slog.Error("Failed to open a channel: " + err.Error())
	}

	err = channel.ExchangeDeclare(
		"mails",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		slog.Error("Failed to declare an exchange: " + err.Error())
	}

	return &Producer[T]{
		QueueName: "mails",
		Channel:   channel,
	}
}

func (p *Producer[T]) Publish(mail T) {
	body := []byte{}

	err := p.Channel.Publish(
		"mails",
		"mail",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		slog.Error("Failed to publish a message: " + err.Error())
	}

	slog.Info("Published a message")
}
