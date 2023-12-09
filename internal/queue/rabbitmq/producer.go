package rabbitmq

import (
	"demo-smtp/internal/config"
	"encoding/json"
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

	queue, err := channel.QueueDeclare(
		"mails", // name
		false,   // durable
		false,   // auto delete
		false,   // exclusive
		false,   // no wait
		nil,     // args
	)

	if err != nil {
		slog.Error("Failed to declare a queue: " + err.Error())
	}

	err = channel.QueueBind(
		queue.Name, // queue name
		"mail",     // routing key
		"mails",    // exchange
		false,      // no wait
		nil,        // args
	)

	if err != nil {
		slog.Error("Failed to bind a queue: " + err.Error())
	}

	return &Producer[T]{
		QueueName: "mails",
		Channel:   channel,
	}
}

func (p *Producer[T]) Publish(mail T) {
	body, err := json.Marshal(mail)

	if err != nil {
		slog.Error("Failed to marshal mail: " + err.Error())
	}

	err = p.Channel.Publish(
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
