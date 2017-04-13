package main

import (
	"fmt"
	"time"

	. "github.com/LFZJun/go-lib/cache/implement"
)

func main() {
	m := NewCache()
	m.Set("name", "ljun", time.Millisecond*500)
	m.Set("name", "ljun", -1)
	time.Sleep(time.Second * 1)
	m.Set("name", "ljun", time.Millisecond*500)
	time.Sleep(time.Second * 1)
	fmt.Println(m.Get("name"))
}
