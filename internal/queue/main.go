package queue

import (
	"demo-smtp/internal/types"
	"fmt"
)

var Queue types.Queue[types.Mail]

func NewQueue() types.Queue[types.Mail] {
	fmt.Println("NewQueue")
	return Queue
}
