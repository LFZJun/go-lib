package adapter

import (
	"testing"
)

func TestDuck(t *testing.T) {
	var duck Duck = &TurkeyAdapter{&WildTurkey{}}
	duck.Quack()
	duck.Fly()
}
