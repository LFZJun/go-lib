package cache

import (
	"context"
	"sync"
	"time"
)

type (
	Store struct {
		Value    interface{}
		timeout  time.Duration
		deadline time.Time

		done   <-chan struct{}
		cancel context.CancelFunc
	}
	MapMutex struct {
		Map   map[string]*Store
		Mutex sync.RWMutex
	}
	CacheTTL struct {
		m *MapMutex
	}
)

func (m *MapMutex) Set(key string, store *Store) {
	m.Mutex.Lock()
	m.Map[key] = store
	m.Mutex.Unlock()
}

func (m *MapMutex) Get(key string) *Store {
	m.Mutex.RLock()
	if v, ok := m.Map[key]; ok {
		m.Mutex.RUnlock()
		return v
	}
	m.Mutex.RUnlock()
	return nil
}

func (m *MapMutex) Del(key string) {
	m.Mutex.Lock()
	delete(m.Map, key)
	m.Mutex.Unlock()
}

func (c *CacheTTL) tick(key string, s *Store) {
	select {
	case <-time.After(s.timeout):
		c.m.Del(key)
	case <-s.done:
	}
}

func (c *CacheTTL) set(key string, s *Store) {
	switch {
	case s.timeout >= 0:
		ctx, cancel := context.WithCancel(context.Background())
		s.done = ctx.Done()
		s.cancel = cancel
		go c.tick(key, s)
	default:
	}
	if vv := c.m.Get(key); vv != nil && vv.timeout >= 0 {
		vv.cancel()
	}
	c.m.Set(key, s)
}

func (c *CacheTTL) SetDeadline(key string, value interface{}, deadline time.Time) {
	c.set(key, &Store{Value: value, timeout: time.Until(deadline), deadline: deadline})
}

func (c *CacheTTL) SetTimeout(key string, value interface{}, timeout time.Duration) {
	c.set(key, &Store{Value: value, timeout: timeout, deadline: time.Now().Add(timeout)})
}

func (c *CacheTTL) Get(key string) *Store {
	return c.m.Get(key)
}

func NewMapMutex() *MapMutex {
	return &MapMutex{Map: make(map[string]*Store)}
}

func NewCacheTTL() *CacheTTL {
	return &CacheTTL{NewMapMutex()}
}
