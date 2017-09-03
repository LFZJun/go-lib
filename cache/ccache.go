package cache

import (
	"github.com/LFZJun/go-lib/cache/timer"
	"sync"
	"time"
)

type (
	Pair struct {
		Task  *timer.Task
		Value interface{}
	}
	CCache struct {
		cMap  *sync.Map
		timer *timer.TimingWheel
	}
)

func (c *CCache) Get(key interface{}) (interface{}, bool) {
	value, ok := c.cMap.Load(key)
	if !ok {
		return nil, ok
	}
	return value.(*Pair).Value, ok
}

func (c *CCache) Set(key interface{}, value interface{}, timeout time.Duration) {
	// 如果有则删除timeout
	if v, ok := c.cMap.Load(key); ok {
		c.timer.Del(v.(*Pair).Task)
	}
	// task
	task := &timer.Task{
		Timeout: timeout,
		Work: func() {
			c.cMap.Delete(key)
		},
	}
	// 存入map
	c.cMap.Store(key, &Pair{
		Task:  task,
		Value: value,
	})
	// 生成timeout
	c.timer.After(task)
}

func (c *CCache) Del(key interface{}) {
	if v, ok := c.cMap.Load(key); ok {
		c.timer.Del(v.(*Pair).Task)
		c.cMap.Delete(key)
	}
}

func (c *CCache) Stop() {
	c.timer.Stop()
}

func NewCCache() *CCache {
	return &CCache{
		cMap:  &sync.Map{},
		timer: timer.NewTimingWheel(10*time.Millisecond, 1<<8),
	}
}
