package main

import (
	"fmt"
	"time"

	. "github.com/LFZJun/go-lib/cache/implement"
)

func main() {
	m := NewMaps()
	m.Set("name", "ljun", time.Second*4)
	m.Set("school", "yinshua", time.Second*1)
	fmt.Println(m.Get("name"), m.Get("school"))
	time.Sleep(time.Second * 2)
	fmt.Println(m.Get("name"), m.Get("school"))
	m.Del("name")
	fmt.Println(m.Get("name"))
}
