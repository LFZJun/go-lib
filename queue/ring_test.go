package queue

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	shared := 2 << 4
	s := NewQueue(shared)
	for i := 0; i <= shared; i++ {
		if err := s.Push(i + 1); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(s)
	for i := 0; i <= shared; i++ {
		if value, err := s.Pop(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}
	fmt.Println(s)
}
