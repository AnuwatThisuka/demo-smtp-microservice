package rabbitmq

import (
	"demo-smtp/internal/config"
	"demo-smtp/internal/types"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/streadway/amqp"
)

type Consumer[T any] struct {
	QueueName string
	Channel   *amqp.Channel
	ReadCh    chan types.ReadData[T]
}

func NewConsumer[T any]() *Consumer[T] {
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

	return &Consumer[T]{
		QueueName: "mails",
		Channel:   channel,
	}
}

func (c *Consumer[T]) Read(ch chan<- types.ReadData[T]) {

	for {
		message, err := c.Channel.Consume(
			c.QueueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			ch <- types.ReadData[T]{Data: nil, Err: err}
			continue
		}

		var model T

		delivery := <-message

		err = json.Unmarshal(delivery.Body, &model)

		if err != nil {
			ch <- types.ReadData[T]{Data: nil, Err: err}
			continue
		}

		ch <- types.ReadData[T]{Data: &model, Err: nil}
	}
}

func (c *Consumer[T]) Close() error {
	return c.Channel.Close()
}
