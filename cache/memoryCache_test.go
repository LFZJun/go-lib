package cache

import (
	"testing"
	"time"
)

func BenchmarkCache_Get(b *testing.B) {
	m := NewCache()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Get(time.Now().String())
		}
	})
}
