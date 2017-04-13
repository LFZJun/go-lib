package implement

import (
	"fmt"
	"time"

	"github.com/LFZJun/go-lib/lock/logic"
	"gopkg.in/redis.v4"
	"qiniupkg.com/x/errors.v7"
)

var (
	ExceedTime = errors.New("获取锁超时")
	DelFail    = errors.New("redis出错")
)

type ReMutex struct {
	client *redis.Client
	Expire time.Duration
	Times  int
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
		fmt.Println(err)
		return false
	}
	return re
}

func (rm *ReMutex) Lock(var1 string) error {
	n := 0
	for !rm.getLock(var1) {
		if n > rm.Times {
			return ExceedTime
		}
		time.Sleep(time.Second)
		n++
	}
	return nil
}

func (rm *ReMutex) Unlock(var1 string) error {
	if rm.ttl(var1) > time.Duration(0) {
		_, err := rm.client.Del(var1).Result()
		if err != nil {
			fmt.Println(err)
			return DelFail
		}
	}
	return nil
}

func NewReMutex(r *Redis) logic.Mutex {
	return &ReMutex{r.Client, time.Minute, 10}
}
