package rabbitmq

import (
	"demo-smtp/internal/config"
	"demo-smtp/internal/types"
	"fmt"
	"log/slog"

	"github.com/streadway/amqp"
)

type Consumer[T any] struct {
	QueueName string
	Channel   *amqp.Channel
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

	q, err := c.Channel.QueueDeclare(
		c.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		slog.Error("Failed to declare a queue: " + err.Error())
	}

	msgs, err := c.Channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		slog.Error("Failed to register a consumer: " + err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			ch <- types.ReadData[T]{
				Data: d.Body,
			}
		}
	}()

	slog.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
