package lock

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
	"os"
	"time"
)

var (
	ExceedTime = errors.New("获取锁超时")
	DelFail    = errors.New("redis出错")
	Logger     = log.New(os.Stderr, "[mutex] ", log.Lshortfile)
)

type ReMutex struct {
	client *redis.Client
	Expire time.Duration
	Retry  int
}

func (rm *ReMutex) ttl(var1 string) time.Duration {
	t, err := rm.client.TTL(var1).Result()
	if err != nil {
		return time.Second
	}
	return t
}

func (rm *ReMutex) getLock(var1 string) bool {
	b := rm.client.SetNX(var1, 1, rm.Expire)
	re, err := b.Result()
	if err != nil {
		Logger.Println(err)
		return false
	}
	return re
}

func (rm *ReMutex) Lock(var1 string) error {
	n := 0
	for !rm.getLock(var1) {
		if n > rm.Retry {
			return ExceedTime
		}
		<-time.After(time.Microsecond)
		n++
	}
	return nil
}

func (rm *ReMutex) Unlock(var1 string) error {
	if rm.ttl(var1) > time.Duration(0) {
		_, err := rm.client.Del(var1).Result()
		if err != nil {
			Logger.Println(err)
			return DelFail
		}
	}
	return nil
}

func NewReMutex(r *Redis, expire time.Duration, retry int) Mutex {
	return &ReMutex{r.Client, expire, retry}
}
