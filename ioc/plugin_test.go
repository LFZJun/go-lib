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
	A string `ioc:"#.abc"`
	B string `ioc:"@.*"`
	C string `ioc:"@.ioc.Bar1"`
}

func TestTomlLoad(t *testing.T) {
	err := TomlLoad(`abc="1"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	Add(new(Bar1))
	Add(new(Foo1))
	AddWithName(new(Foo1), "Foo1")
	Start()
	spew.Dump(GetWithName("ioc.Foo1"))
	spew.Dump(GetWithName("Foo1"))
}
