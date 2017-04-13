package logic

// factory
type Store interface {
	CreatePizza() Pizza
}

// proxy
type PizzaStore struct {
	Store Store
}

func (ps *PizzaStore) OrderPizza() {
	pizza := ps.Store.CreatePizza()
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}
