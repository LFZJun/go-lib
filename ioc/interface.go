package ioc

type Stone interface{}

type Init interface {
	Init()
}

type Ready interface {
	Ready()
}
