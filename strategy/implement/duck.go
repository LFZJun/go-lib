package implement

import (
	"github.com/LFZJun/go-lib/strategy/logic"
)

//duck
func NewDuck(quack logic.QuackBehavior) *logic.Duck {
	return &logic.Duck{QuackInterface: quack}
}
