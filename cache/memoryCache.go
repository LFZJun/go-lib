package cache

import (
	"sync"
	"time"
)

type Value struct {
	V      interface{}
	expire time.Duration
	done   chan struct{}
}

type Maps struct {
	Map   map[string]*Value
	Mutex sync.RWMutex
}

func (m *Maps) tick(key string, value *Value) {
	select {
	case <-time.After(value.expire):
		m.Mutex.Lock()
		delete(m.Map, key)
		m.Mutex.Unlock()
	case <-value.done:
	}
}

func (m *Maps) Set(key string, value interface{}, expire time.Duration) {
	var v *Value
	switch {
	case expire >= 0:
		v = &Value{value, expire, make(chan struct{})}
	default:
		v = &Value{value, expire, nil}
	}
	m.Mutex.Lock()
	vv, ok := m.Map[key]
	if ok && vv.expire >= 0 {
		vv.done <- struct{}{}
	}
	m.Map[key] = v
	m.Mutex.Unlock()
	if v.expire >= 0 {
		go m.tick(key, v)
	}

}

func (m *Maps) Get(key string) interface{} {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	v, ok := m.Map[key]
	if !ok {
		return nil
	}
	return v.V
}

func (m *Maps) Del(key string) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	v, ok := m.Map[key]
	if !ok {
		return nil
	}
	if v.expire >= 0 {
		v.done <- struct{}{}
	}
	delete(m.Map, key)
	return nil
}

type Cache interface {
	Set(key string, value interface{}, expire time.Duration)
	Get(key string) interface{}
	Del(key string) error
}

func NewCache() Cache {
	return &Maps{Map: make(map[string]*Value)}
}
