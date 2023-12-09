package rabbitmq

import (
	"context"
	"demo-smtp/internal/types"
)

type RabbitQueue[T any] struct {
	producer *Producer[T]
}

// Ping implements types.Queue.
func (*RabbitQueue[T]) Ping() (string, error) {
	panic("unimplemented")
}

// Read implements types.Queue.
func (*RabbitQueue[T]) Read(ctx context.Context) {
	panic("unimplemented")
}

func (*RabbitQueue[T]) Write(mail types.Mail) error {
	panic("unimplemented")
}

// NewRabbitQueue returns a new RabbitQueue.
func NewRabbitQueue() types.Queue[types.Mail] {
	return &RabbitQueue[types.Mail]{
		producer: NewProducer[types.Mail](),
	}
}
