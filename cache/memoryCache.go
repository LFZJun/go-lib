package cache

import (
	"context"
	"sync"
	"time"

	"github.com/LFZJun/go-lib/cache/hashl"
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
	TTL struct {
		m *MapMutex
	}
	TTLCache struct {
		size int
		ttls []*TTL
	}
)

func (m *MapMutex) Set(key string, store *Store) {
	m.Mutex.Lock()
	m.Map[key] = store
	m.Mutex.Unlock()
}

func (m *MapMutex) Get(key string) *Store {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	return m.Map[key]
}

func (m *MapMutex) Del(key string) {
	m.Mutex.Lock()
	delete(m.Map, key)
	m.Mutex.Unlock()
}

func (t *TTL) tick(key string, s *Store) {
	select {
	case <-time.After(s.timeout):
		t.m.Del(key)
	case <-s.done:
	}
}

func (t *TTL) set(key string, s *Store) {
	switch {
	case s.timeout >= 0:
		ctx, cancel := context.WithCancel(context.Background())
		s.done = ctx.Done()
		s.cancel = cancel
		go t.tick(key, s)
	default:
	}
	if vv := t.m.Get(key); vv != nil && vv.timeout >= 0 {
		vv.cancel()
	}
	t.m.Set(key, s)
}

func (t *TTL) SetDeadline(key string, value interface{}, deadline time.Time) {
	t.set(key, &Store{Value: value, timeout: time.Until(deadline), deadline: deadline})
}

func (t *TTL) SetTimeout(key string, value interface{}, timeout time.Duration) {
	t.set(key, &Store{Value: value, timeout: timeout, deadline: time.Now().Add(timeout)})
}

func (t *TTL) Get(key string) *Store {
	return t.m.Get(key)
}

func (tc *TTLCache) Get(key string) *Store {
	index := hashl.HashIndex32(key, uint32(tc.size))
	return tc.ttls[index].Get(key)
}

func (tc *TTLCache) SetDeadline(key string, value interface{}, deadline time.Time) {
	index := hashl.HashIndex32(key, uint32(tc.size))
	tc.ttls[index].SetDeadline(key, value, deadline)
}

func (tc *TTLCache) SetTimeout(key string, value interface{}, timeout time.Duration) {
	index := hashl.HashIndex32(key, uint32(tc.size))
	tc.ttls[index].SetTimeout(key, value, timeout)
}

func NewMapMutex() *MapMutex {
	return &MapMutex{Map: make(map[string]*Store)}
}

func NewTTL() *TTL {
	return &TTL{NewMapMutex()}
}

func NewTTLCache(size int) *TTLCache {
	t := make([]*TTL, 0, size)
	for i := 0; i < size; i++ {
		t = append(t, NewTTL())
	}
	return &TTLCache{size: size, ttls: t}
}
