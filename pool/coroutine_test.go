package pool

import (
	"fmt"
	"testing"
	"time"
	"runtime"
)

func TestNewCoroutinePool(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // 2
	p := NewCoroutinePool(10)
	for i := 0; i < 10; i++ {
		p.Add(func() {
			time.Sleep(time.Second)
			fmt.Println(1)
		})
	}
	fmt.Println(runtime.NumGoroutine()) // 2 + 10 + 1
	p.Close()
	fmt.Println("ok")
	fmt.Println(runtime.NumGoroutine())
}
