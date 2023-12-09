package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Consumer[T any] struct {
	QueueName string
	Channel   *amqp.Channel
	Messages  <-chan amqp.Delivery
}

func NewConsumer[T any](queueName string, channel *amqp.Channel) (*Consumer[T], error) {
	messages, err := channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Consumer[T]{
		QueueName: queueName,
		Channel:   channel,
		Messages:  messages,
	}, nil
}
