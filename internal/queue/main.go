package queue

import (
	"context"
	"demo-smtp/internal/queue/rabbitmq"
	"demo-smtp/internal/types"
)

var Queue types.Queue[types.Mail]
var cancel context.CancelFunc

func NewQueue() types.Queue[types.Mail] {
	Queue = rabbitmq.NewRabbitQueue()
	return Queue
}

func StartWorker(queue types.Queue[types.Mail]) {
	var ctx context.Context

	ctx, cancel = context.WithCancel(context.Background())
	go queue.Read(ctx)
}

func StopWorker() {
	cancel()
}
