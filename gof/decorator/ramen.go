package decorator

type Ramen struct{}

func (p Ramen) Description() string {
	return "ramen"
}

func (p Ramen) Price() float32 {
	return 10
}
