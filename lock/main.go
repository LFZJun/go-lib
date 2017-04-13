package main

import (
	"fmt"
	"sync"

	"github.com/LFZJun/go-lib/lock/implement"
)

func main() {
	mutex := implement.NewReMutex(implement.NewRedis(implement.RedisUser{Address: "127.0.0.1:6379"}))
	group := sync.WaitGroup{}
	mutex.Unlock("ljun")
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			err := mutex.Lock("ljun")
			if err != nil {
				fmt.Println(err)
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
