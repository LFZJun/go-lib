package strategy

import (
	"fmt"
)

// Quack
type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack")
}
