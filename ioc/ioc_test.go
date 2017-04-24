package ioc

import (
	"testing"
)

type Bar struct {
	Value string
}
type Foo struct {
	Name  *Bar `ioc:"*"`
	Name1 *Bar `ioc:"a.b"`
 	Sex   string
}

func TestPut(t *testing.T) {
	Put(&Bar{
		Value: "dc",
	})
	foo := new(Foo)
	Put(foo)
	Start()
	//spew.Dump(foo)
	//spew.Dump(GetStoneWithName("Bar"))
}
