package main

import (
	"sync"
	"testing"
	"time"
	"fmt"
)

func TestName(t *testing.T) {
	mutex := NewReMutex(NewRedis(RedisUser{Address: "127.0.0.1:6379"}), time.Minute, 100)
	group := sync.WaitGroup{}
	mutex.Unlock("ljun")
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			err := mutex.Lock("ljun")
			if err != nil {
				Logger.Fatalln(err)
				group.Done()
				return
			}
			defer mutex.Unlock("ljun")
			fmt.Println("dc")
			group.Done()
		}()
	}
	group.Wait()
}
