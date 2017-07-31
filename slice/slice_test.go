package slice

import (
	"fmt"
	"testing"
)

type foo struct {
	Id    int
	Value int64
}

func TestGroupBy(t *testing.T) {
	srcFoo := []foo{{1, 1}, {1, 3}, {1, 4}}
	hashFunc := func(h interface{}) interface{} {
		return h.(foo).Id
	}
	err := GroupBy(&srcFoo, hashFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(srcFoo)
}

func TestFilter(t *testing.T) {
	srcFoo := []foo{{1, 1}, {1, 3}, {1, 4}}
	filterFunc := func(i interface{}) bool {
		t := i.(foo)
		return t.Value < 4
	}
	err := Filter(&srcFoo, filterFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(srcFoo)
}
