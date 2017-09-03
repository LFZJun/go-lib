package cache

import (
	"testing"
	"sync"
	"time"
)

func BenchmarkSyncMap(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			m.Store(now, now)
		}
	})
}

func BenchmarkCCache_Set(b *testing.B) {
	c := NewCCache()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			c.Set(now, time.Now().Add(time.Second*10), time.Second*10)
		}
	})
}
