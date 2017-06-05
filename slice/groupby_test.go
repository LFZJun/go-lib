package slice

import (
	"fmt"
	"testing"
)

type Foo struct {
	Id    int
	Value int64
}

func TestGroupBy(t *testing.T) {
	srcFoo := []Foo{{1, 1}, {1, 3}, {1, 4}}
	destFoo := []Foo{}
	hashFunc := func(h interface{}) interface{} {
		return h.(Foo).Id
	}
	cmpFunc := func(i interface{}, j interface{}) bool {
		return i.(Foo).Value < j.(Foo).Value
	}
	err := GroupBy(&destFoo, srcFoo, hashFunc, cmpFunc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(destFoo)
}
