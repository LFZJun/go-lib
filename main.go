package main

import (
	"github.com/LFZJun/go-lib/ioc"
	"github.com/davecgh/go-spew/spew"
)

type Bar struct {
	Value string
}

type Foo struct {
	Name *Bar `ioc:"*"`
	Sex  string
}

func main() {
	ioc.Put(&Bar{
		Value: "dc",
	})
	foo := new(Foo)
	ioc.Put(foo)
	ioc.Start()
	spew.Dump(foo)
}
