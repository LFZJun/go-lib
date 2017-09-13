package ioc

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

type (
	FooBaz interface {
		Names()
	}

	Bar struct {
		Value string
	}

	Foo struct {
		Name  *Bar    `ioc:"*"`
		Name1 *Bar    `ioc:"a.b"`
		Name2 FooBaz  `ioc:"*"`
		Sex   *string `ioc:"*"`
	}
)

func (f *Bar) Names() {

}

func TestPut(t *testing.T) {
	foo := new(Foo)
	sex := "sex"
	Put(foo)
	Put(&Bar{Value: "dc"})
	Put(&sex)
	Start()
	spew.Dump(foo)
}
