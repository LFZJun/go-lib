package ioc

type (
	Stone interface{}
	Init  interface {
		Init()
	}
	Ready interface {
		Ready()
	}
)
