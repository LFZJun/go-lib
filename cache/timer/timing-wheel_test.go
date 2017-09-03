package timer

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestNewTimingWheel(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(1)
	timing := NewTimingWheel(time.Millisecond*10, 128)
	start := time.Now()
	timing.After(
		NewTask(time.Millisecond*1500,
			func() {
				end := time.Now()
				fmt.Println(end.Sub(start))
				fmt.Println(runtime.NumGoroutine())
				wait.Done()
			}))
	wait.Wait()
}
