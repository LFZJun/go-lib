package decorator

type Egg struct {
	Noddles Food
}

func (p Egg) SetNoddles(noddles Food) {
	p.Noddles = noddles
}

func (p Egg) Description() string {
	return p.Noddles.Description() + "+egg"
}

func (p Egg) Price() float32 {
	return p.Noddles.Price() + 2
}
