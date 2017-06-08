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
	destFoo := []foo{}
	hashFunc := func(h interface{}) interface{} {
		return h.(foo).Id
	}
	cmpFunc := func(i interface{}, j interface{}) bool {
		return i.(foo).Value < j.(foo).Value
	}
	err := GroupBy(&destFoo, srcFoo, hashFunc, cmpFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(destFoo)
}

func TestFilter(t *testing.T) {
	srcFoo := []foo{{1, 1}, {1, 3}, {1, 4}}
	filterFunc := func(i interface{}) bool {
		t := i.(foo)
		return t.Value == 1
	}
	err := Filter(&srcFoo, filterFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(srcFoo)
}
