package pool

import (
	"sync"
)

type (
	Task func()

	CoroutinePool interface {
		Add(Task)
		Wait()
		Close()
	}

	Buf struct {
		queue []Task
		lock  sync.Mutex
	}

	coroutinePool struct {
		wait sync.WaitGroup
		auto chan struct{}
		buf  *Buf
		Size int
		Task chan Task
	}
)

func (b *Buf) Push(task Task) {
	b.lock.Lock()
	b.queue = append(b.queue, task)
	b.lock.Unlock()
}

func (b *Buf) Pop() (Task, bool) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if len(b.queue) == 0 {
		return nil, false
	}
	t := b.queue[0]
	b.queue = b.queue[1:]
	return t, true
}

func (c *coroutinePool) coroutineRun() {
	for {
		fn, ok := <-c.Task
		if !ok {
			return
		}
		fn()
		c.auto <- struct{}{}
		c.wait.Add(-1)
	}
}

func (c *coroutinePool) autoAdd() {
	for {
		_, ok := <-c.auto
		if !ok {
			return
		}
		if t, has := c.buf.Pop(); has {
			c.Task <- t
		}
	}
}

func (c *coroutinePool) Add(fn Task) {
	c.wait.Add(1)
	if len(c.Task) < c.Size {
		c.Task <- fn
		return
	}
	// 多余的任务放入缓冲区
	c.buf.Push(fn)
}

func (c *coroutinePool) Wait() {
	c.wait.Wait()
}

func (c *coroutinePool) Close() {
	// 等待任务执行完毕
	c.Wait()
	// 停止自动添加任务
	close(c.auto)
	// 停止池中所有协程
	close(c.Task)
}

func initCoroutinePool(pool *coroutinePool) {
	pool.buf = &Buf{
		queue: make([]Task, 0, pool.Size),
	}
	pool.auto = make(chan struct{}, 1)
	for i := 0; i < pool.Size; i++ {
		// 创建goroutine
		go pool.coroutineRun()
	}
	// 用一个goroutine维持任务分发
	go pool.autoAdd()
}

func NewCoroutinePool(size int) CoroutinePool {
	if size < 1 {
		size = 1
	}
	c := &coroutinePool{
		Task: make(chan Task, size),
		Size: size,
	}
	initCoroutinePool(c)
	return c
}
