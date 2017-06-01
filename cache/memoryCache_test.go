package cache

import (
	"testing"
	"time"
)

func BenchmarkCache_Get(b *testing.B) {
	c := NewTTLCache(1 << 8)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get(time.Now().String())
		}
	})
}

func BenchmarkTTLCache_SetDeadline(b *testing.B) {
	c := NewTTLCache(1 << 8)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			c.SetDeadline(now, now, time.Now().Add(time.Second*10))
		}
	})
}
