package ioc

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type Bar1 struct {
}

func (b *Bar1) BeforeInit() interface{} {
	return "2"
}

type Foo1 struct {
	A string `ioc:"#.ab"`
	B string `ioc:"@.*"`
	C string `ioc:"@.ioc.Bar1"`
}

func TestTomlLoad(t *testing.T) {
	err := TomlLoad(`ab="1"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	Put(new(Bar1))
	Put(new(Foo1))
	PutWithName(new(Foo1), "Foo1")
	Start()
	spew.Dump(GetWithName("ioc.Foo1"))
	spew.Dump(GetWithName("Foo1"))
}
