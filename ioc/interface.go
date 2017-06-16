package ioc

type (
	Water interface{}
	Init  interface {
		Init()
	}
	Ready interface {
		Ready()
	}
)
