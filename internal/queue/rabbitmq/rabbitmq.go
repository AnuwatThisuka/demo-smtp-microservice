package rabbitmq

import (
	"context"
	"demo-smtp/internal/smtp"
	"demo-smtp/internal/types"
	"fmt"
	"log/slog"
)

// RabbitQueue is a RabbitMQ implementation of types.Queue.
type RabbitQueue[T any] struct {
	producer *Producer[T]
	consumer *Consumer[types.Mail]
}

// Ping implements types.Queue.
func (*RabbitQueue[T]) Ping() (string, error) {
	panic("Ping unimplemented")
}

// Read implements types.Queue.
func (r *RabbitQueue[T]) Read(ctx context.Context) {
	slog.Info("RabbitQueue.Read")

	readCh := make(chan types.ReadData[types.Mail])

	go r.consumer.Read(readCh)

	for {
		select {
		case <-ctx.Done():
			slog.Info("Stopping RabbitQueue consumer...")

			_ = r.consumer.Close()
			return
		case data := <-readCh:
			if data.Err != nil {
				slog.Error("Failed to read content", data.Err)
				continue
			}

			fmt.Println(data)
			err := smtp.SendEmail(data.Data)
			slog.Info("Sending email")

			if err != nil {
				slog.Error("Failed to send email", err)
				continue
			}

			slog.Info("Sent email successfully")
		}
	}
}

// Write implements types.Queue.
func (r *RabbitQueue[T]) Write(mail T) error {
	r.producer.Publish(mail)
	return nil
}

// NewRabbitQueue returns a new RabbitQueue.
func NewRabbitQueue() types.Queue[types.Mail] {
	return &RabbitQueue[types.Mail]{
		producer: NewProducer[types.Mail](),
		consumer: NewConsumer[types.Mail](),
	}
}
