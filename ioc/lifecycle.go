package ioc

type Lifecycle int

const (
	BeforeInit = iota
	BeforeReady
)
