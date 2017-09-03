package cache

import (
	"testing"
	"time"
)

func BenchmarkTTLCache_SetDeadline(b *testing.B) {
	c := NewTTLCache(1 << 8)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			c.SetDeadline(now, now, time.Now().Add(time.Second*10))
		}
	})
}

func BenchmarkTTL_SetDeadline(b *testing.B) {
	t := NewTTL()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			t.SetDeadline(now, now, time.Now().Add(time.Second*10))
		}
	})
}

func BenchmarkTTLCache_Set(b *testing.B) {
	c := NewTTLCache(1 << 8)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now().String()
			c.Set(now, &Store{
				Timeout: time.Second * 10,
				Value:   time.Now().Add(time.Second * 10),
			})
		}
	})
}
