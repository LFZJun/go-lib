package ioc

import (
	"fmt"
	"testing"
)

type TestPl struct {
}

func (p *TestPl) Value(path string) interface{} {
	return "a"
}

func (p *TestPl) Prefix() string {
	return "#"
}

func (p *TestPl) Priority() int {
	return 0
}

type Foo1 struct {
	Name string `ioc:"#.name"`
}

func TestPlugin(t *testing.T) {
	RegisterPlugin(BeforeInit, new(TestPl))
	Put(new(Foo1))
	Start()
	fmt.Println(GetStoneWithName("Foo1"))
}
