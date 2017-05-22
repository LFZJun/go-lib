package singleton

var Foo *foo = New()

type foo struct {
}

func (f *foo) Bar() string {
	return "bar"
}

func New() *foo {
	return new(foo)
}
