package cache

import (
	"testing"
	"time"
)

func BenchmarkCache_Get(b *testing.B) {
	c := NewCacheTTL()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Get(time.Now().String())
		}
	})
}

func BenchmarkName(b *testing.B) {
	c := NewCacheTTL()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			c.SetDeadline(now, now, time.Now().Add(time.Second*10))
		}
	})
}
