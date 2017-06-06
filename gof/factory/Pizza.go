package factory

// worker
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
