package queue

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	s := NewQueue(5)
	for i := 0; i <= 5; i++ {
		if err := s.Push(i + 1); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(s)
	for i := 0; i <= 5; i++ {
		if value, err := s.Pop(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}
	fmt.Println(s)
}
