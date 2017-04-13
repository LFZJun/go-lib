package implement

import (
	"github.com/LFZJun/go-lib/decorator/logic"
)

type Egg struct {
	Noddles logic.Food
}

func (p Egg) SetNoddles(noddles logic.Food) {
	p.Noddles = noddles
}

func (p Egg) Description() string {
	return p.Noddles.Description() + "+egg"
}

func (p Egg) Price() float32 {
	return p.Noddles.Price() + 2
}
