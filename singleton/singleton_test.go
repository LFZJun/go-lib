package singleton

import "testing"

func TestSingleton(t *testing.T) {
	foo := Foo
	foo.Bar()
}
