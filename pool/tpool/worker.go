package tpool

import "time"

type (
	Future struct {
	}

	Command func() error
	Task    func() Future
)

type (
	Queue interface {
		Offer(interface{}) bool
		Remove(interface{}) bool
		IsEmpty() bool
	}

	Worker struct {
		tunnel      chan Command
		lastUseTime time.Time
	}
)

func (w *Worker) Start() {
	for {
		t, ok := <-w.tunnel
		if !ok {
			return
		}
		t()
	}
}

func (w *Worker) Put(c Command) {
	w.tunnel <- c
}

func (w *Worker) Close() {
	close(w.tunnel)
}
