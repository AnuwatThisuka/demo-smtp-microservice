package rabbitmq

import (
	"context"
	"fmt"
)

type RabbitQueue[T any] struct {
}

func (r *RabbitQueue[T]) Read(ctx context.Context) {
}

func (r *RabbitQueue[T]) Write(data T) {
	fmt.Println("RabbitMQ: Write", data)
}
