package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	m := NewCache()
	m.Set("name", "ljun", time.Millisecond*500)
	m.Set("name", "ljun", -1)
	fmt.Println(m.Get("name"))
	time.Sleep(time.Second * 1)
	m.Set("name", "ljun", time.Millisecond*500)
	time.Sleep(time.Second * 1)
	fmt.Println(m.Get("name"))
}
