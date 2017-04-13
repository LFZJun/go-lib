package logic

// worker
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
