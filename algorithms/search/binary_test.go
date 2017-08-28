package search

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	foo := []int{4}
	fmt.Println(BinarySearch(foo, 7))
}
