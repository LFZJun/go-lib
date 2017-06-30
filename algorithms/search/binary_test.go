package search

import (
	"testing"
	"fmt"
)

func TestBinary(t *testing.T) {
	foo := []int{4}
	fmt.Println(BinarySearch(foo, 7))
}
