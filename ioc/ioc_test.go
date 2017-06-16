package ioc

import (
	"testing"
)

type (
	Bar struct {
		Value string
	}

	Foo struct {
		Name  *Bar `ioc:"*"`
		Name1 *Bar `ioc:"a.b"`
		Sex   string
	}
)

func TestPut(t *testing.T) {
	Add(&Bar{
		Value: "dc",
	})
	foo := new(Foo)
	Add(foo)
	Start()
	//spew.Dump(foo)
	//spew.Dump(GetWithName("Bar"))
}
