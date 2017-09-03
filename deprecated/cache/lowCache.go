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
	SafeMap struct {
		Map   map[string]*Store
		Mutex sync.RWMutex
	}
	Segment struct {
		m *SafeMap
	}
	TTLCache struct {
		size     uint32
		segments []*Segment
	}
)

func (m *SafeMap) Set(key string, store *Store) {
	m.Mutex.Lock()
	m.Map[key] = store
	m.Mutex.Unlock()
}

func (m *SafeMap) Get(key string) *Store {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	return m.Map[key]
}

func (m *SafeMap) Del(key string) {
	m.Mutex.Lock()
	delete(m.Map, key)
	m.Mutex.Unlock()
}

func (t *Segment) tick(key string, s *Store) {
	select {
	case <-time.After(s.Timeout):
		t.m.Del(key)
	case <-s.done:
	}
}

func (t *Segment) Set(key string, s *Store) {
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

func (t *Segment) SetDeadline(key string, value interface{}, deadline time.Time) {
	t.Set(key, &Store{Value: value, Timeout: time.Until(deadline), Deadline: deadline})
}

func (t *Segment) SetTimeout(key string, value interface{}, timeout time.Duration) {
	t.Set(key, &Store{Value: value, Timeout: timeout, Deadline: time.Now().Add(timeout)})
}

func (t *Segment) Get(key string) *Store {
	return t.m.Get(key)
}

func (tc *TTLCache) Get(key string) *Store {
	index := hashl.HashIndex32(key, tc.size)
	return tc.segments[index].Get(key)
}

func (tc *TTLCache) Set(key string, s *Store) {
	index := hashl.HashIndex32(key, tc.size)
	tc.segments[index].Set(key, s)
}

func (tc *TTLCache) SetDeadline(key string, value interface{}, deadline time.Time) {
	index := hashl.HashIndex32(key, tc.size)
	tc.segments[index].SetDeadline(key, value, deadline)
}

func (tc *TTLCache) SetTimeout(key string, value interface{}, timeout time.Duration) {
	index := hashl.HashIndex32(key, tc.size)
	tc.segments[index].SetTimeout(key, value, timeout)
}

// deprecated
func NewMapMutex() *SafeMap {
	return &SafeMap{Map: make(map[string]*Store)}
}

// deprecated
func NewTTL() *Segment {
	return &Segment{NewMapMutex()}
}

// deprecated
func NewTTLCache(size uint32) *TTLCache {
	t := make([]*Segment, 0, size)
	for i := 0; i < int(size); i++ {
		t = append(t, NewTTL())
	}
	return &TTLCache{size: size, segments: t}
}
