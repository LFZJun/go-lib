package main

import (
	"github.com/LFZJun/go-lib/singleton/implement"
)

func main() {
	foo := implement.Foo
	foo.Bar()
}
