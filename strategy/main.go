package main

import (
	"github.com/LFZJun/go-lib/strategy/implement"
)

func main() {
	duck := implement.NewDuck(new(implement.Quack))
	duck.PerformQuack()
}
