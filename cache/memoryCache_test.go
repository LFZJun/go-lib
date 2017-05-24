package cache

import (
	"log"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/cocotyty/cache"
)

//func BenchmarkCache_Set(b *testing.B) {
//	m := NewCache()
//
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			m.Set(time.Now().String(), time.Now().String(), time.Second)
//		}
//	})
//}

func BenchmarkCMap_Set(b *testing.B) {
	cmap := cache.NewCMap(32)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cmap.Set(time.Now().String(), &cache.CacheValue{Exp: int64(time.Second), Value: time.Now().String()})
		}
	})
}

func BenchmarkBigCache_Set(b *testing.B) {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 32,
		// time after which entry can be evicted
		LifeWindow: 10 * time.Minute,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// prints information about additional memory allocation
		Verbose: true,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,
		// callback fired when the oldest entry is removed because of its
		// expiration time or no space left for the new entry. Default value is nil which
		// means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,
	}

	c, initErr := bigcache.NewBigCache(config)
	if initErr != nil {
		log.Fatal(initErr)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set(time.Now().String(), []byte(time.Now().String()))
		}
	})
}
