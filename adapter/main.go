package main

import (
	"github.com/LFZJun/go-lib/adapter/implement"
	"github.com/LFZJun/go-lib/adapter/logic"
)

func main() {
	var duck logic.Duck = &implement.TurkeyAdapter{&implement.WildTurkey{}}
	duck.Quack()
	duck.Fly()
}
