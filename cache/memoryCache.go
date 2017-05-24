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

func (v *Value) CancelDel() {
	if v.expire > 0 {
		v.done <- struct{}{}
	}
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
		defer func() {go m.tick(key, v)}()
	default:
		v = &Value{value, expire, nil}
	}
	m.Mutex.Lock()
	if vv, ok := m.Map[key]; ok {
		vv.CancelDel()
	}
	m.Map[key] = v
	m.Mutex.Unlock()
}

func (m *Maps) Get(key string) interface{} {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	if v, ok := m.Map[key]; ok {
		return v.V
	}
	return nil
}

func (m *Maps) Del(key string) error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if v, ok := m.Map[key]; ok {
		v.CancelDel()
		delete(m.Map, key)
	}
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
