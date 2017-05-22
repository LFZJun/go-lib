package decorator

type Sausage struct {
	Noddles Food
}

func (p Sausage) SetNoddles(noddles Food) {
	p.Noddles = noddles
}

func (p Sausage) Description() string {
	return p.Noddles.Description() + "+sausage"
}

func (p Sausage) Price() float32 {
	return p.Noddles.Price() + 2
}
