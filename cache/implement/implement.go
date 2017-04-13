package implement

import (
	"sync"
	"time"
)

type Maps struct {
	Map   map[string]*Value
	Mutex sync.RWMutex
}

func (m *Maps) tick(key string, value *Value) {
	select {
	case <-time.After(value.Expire):
		m.Mutex.Lock()
		delete(m.Map, key)
		m.Mutex.Unlock()
	case <-value.Done:
		return
	}
}

func (m *Maps) Set(key string, value interface{}, expire time.Duration) {
	v := &Value{value, expire, make(chan struct{})}
	m.Mutex.Lock()
	m.Map[key] = v
	m.Mutex.Unlock()
	go m.tick(key, v)

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
	v.Done <- struct{}{}
	delete(m.Map, key)
	return nil
}

type Value struct {
	V      interface{}
	Expire time.Duration
	Done   chan struct{}
}

func NewMaps() *Maps {
	return &Maps{Map: make(map[string]*Value)}
}
