package ioc

import (
	"fmt"
	"testing"
	"github.com/davecgh/go-spew/spew"
)

type Bar struct {
	Value string
}

func (b *Bar) Init() {
	fmt.Println(1)
	b.Value = "yb"
}

func (b *Bar) Ready() {
	fmt.Println(2)
}

type Foo struct {
	Name *Bar `ioc:"*"`
	Sex  string
}

func TestIoc(t *testing.T) {
	Put(&Bar{
		Value: "dc",
	})
	foo := new(Foo)
	Put(foo)
	Start()
	spew.Dump(foo)
}
