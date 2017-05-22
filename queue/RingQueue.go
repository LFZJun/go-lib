package queue

import (
	"errors"
)

type RingQueue interface {
	Push(x interface{}) error
	Pop() (interface{}, error)
}

var (
	ErrQueueFull  = errors.New("queue full.")
	ErrQueueEmpty = errors.New("queue empty.")
)

type ringQueue struct {
	data []interface{}
	head int // 读取位
	tail int // 写入位
	tag  int // 标识位
}

func NewQueue(cap int) RingQueue {
	return &ringQueue{
		data: make([]interface{}, cap),
	}
}

func (q *ringQueue) Push(x interface{}) error {
	if q.head == q.tail && q.tag == 1 {
		return ErrQueueFull
	}
	q.data[q.tail] = x
	q.tail = (q.tail + 1) % cap(q.data)
	if q.tail == q.head {
		q.tag = 1
	}
	return nil
}

func (q *ringQueue) Pop() (interface{}, error) {
	if q.tail == q.head && q.tag == 0 {
		return 0, ErrQueueEmpty
	}
	x := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) % cap(q.data)
	if q.head == q.tail {
		q.tag = 0
	}
	return x, nil
}
