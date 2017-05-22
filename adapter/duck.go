package adapter

type Duck interface {
	Quack()
	Fly()
}

type Turkey interface {
	Gobble()
	Fly()
}