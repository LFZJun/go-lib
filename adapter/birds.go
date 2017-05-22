package adapter

import (
	"fmt"
)

type WildTurkey struct {
}

func (w *WildTurkey) Gobble() {
	fmt.Println("wild Turky Gobble")
}

func (w *WildTurkey) Fly() {
	fmt.Println("wild Turky Fly")
}

type TurkeyAdapter struct {
	Turkey Turkey
}

func (t *TurkeyAdapter) Quack() {
	t.Turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	t.Turkey.Fly()
}
