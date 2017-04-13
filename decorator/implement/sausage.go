package implement

import (
	"github.com/LFZJun/go-lib/decorator/logic"
)

type Sausage struct {
	Noddles logic.Food
}

func (p Sausage) SetNoddles(noddles logic.Food) {
	p.Noddles = noddles
}

func (p Sausage) Description() string {
	return p.Noddles.Description() + "+sausage"
}

func (p Sausage) Price() float32 {
	return p.Noddles.Price() + 2
}
