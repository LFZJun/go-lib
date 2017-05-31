package ioc

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
)

type Foo1 struct {
	A string `ioc:"#.abc"`
}

func TestPlugin(t *testing.T) {
	TomlLoad(`abc="1"`)
	Put(new(Foo1))
	PutWithName(new(Foo1), "Foo1")
	Start()
	spew.Dump(GetStoneWithName("ioc.Foo1"))
	spew.Dump(GetStoneWithName("Foo1"))
}
