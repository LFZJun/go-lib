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
	keyFunc := func(k interface{}) int64 {
		return k.(Foo).Value
	}
	err := GroupBy(&destFoo, srcFoo, hashFunc, keyFunc, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(destFoo)
	destFoo = []Foo{}
	err = GroupBy(&destFoo, srcFoo, hashFunc, keyFunc, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(destFoo)
}
