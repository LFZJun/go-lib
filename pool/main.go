package main

import (
	"fmt"
	"strconv"

	"github.com/LFZJun/go-lib/pool/implement"
)

func main() {
	pool := implement.NewConnectionPool(func() (interface{}, error) {
		return []string{}, nil
	}, 10)
	for i := 0; i < 20; i++ {
		p := pool.Get().([]string)
		p = append(p, strconv.Itoa(i))
		fmt.Println(p)
		pool.Release(p)
	}
}
