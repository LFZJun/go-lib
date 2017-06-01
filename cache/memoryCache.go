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
		Timeout  time.Duration
		Deadline time.Time

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
		size uint32
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
	case <-time.After(s.Timeout):
		t.m.Del(key)
	case <-s.done:
	}
}

func (t *TTL) Set(key string, s *Store) {
	switch {
	case s.Timeout >= 0:
		ctx, cancel := context.WithCancel(context.Background())
		s.done = ctx.Done()
		s.cancel = cancel
		go t.tick(key, s) // 随着 coroutine的数量上升，性能会急剧下降，所以说赶紧换个定时器吧
	default:
	}
	if vv := t.m.Get(key); vv != nil && vv.Timeout >= 0 { // 会有一次读的性能损耗 采用这种定时策略
		vv.cancel()
	}
	t.m.Set(key, s)
}

func (t *TTL) SetDeadline(key string, value interface{}, deadline time.Time) {
	t.Set(key, &Store{Value: value, Timeout: time.Until(deadline), Deadline: deadline})
}

func (t *TTL) SetTimeout(key string, value interface{}, timeout time.Duration) {
	t.Set(key, &Store{Value: value, Timeout: timeout, Deadline: time.Now().Add(timeout)})
}

func (t *TTL) Get(key string) *Store {
	return t.m.Get(key)
}

func (tc *TTLCache) Get(key string) *Store {
	index := hashl.HashIndex32(key, tc.size)
	return tc.ttls[index].Get(key)
}

func (tc *TTLCache) Set(key string, s *Store) {
	index := hashl.HashIndex32(key, tc.size)
	tc.ttls[index].Set(key, s)
}

func (tc *TTLCache) SetDeadline(key string, value interface{}, deadline time.Time) {
	index := hashl.HashIndex32(key, tc.size)
	tc.ttls[index].SetDeadline(key, value, deadline)
}

func (tc *TTLCache) SetTimeout(key string, value interface{}, timeout time.Duration) {
	index := hashl.HashIndex32(key, tc.size)
	tc.ttls[index].SetTimeout(key, value, timeout)
}

func NewMapMutex() *MapMutex {
	return &MapMutex{Map: make(map[string]*Store)}
}

func NewTTL() *TTL {
	return &TTL{NewMapMutex()}
}

func NewTTLCache(size uint32) *TTLCache {
	t := make([]*TTL, 0, size)
	for i := 0; i < int(size); i++ {
		t = append(t, NewTTL())
	}
	return &TTLCache{size: size, ttls: t}
}
