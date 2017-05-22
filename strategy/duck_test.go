package strategy

import "testing"

func TestDuck_PerformQuack(t *testing.T) {
	duck := NewDuck(new(Quack))
	duck.PerformQuack()
}
