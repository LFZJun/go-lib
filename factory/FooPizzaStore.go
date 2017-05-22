package factory

// factory implement
type fooStore struct{}

func (ps *fooStore) CreatePizza() Pizza {
	return new(FooPizza)
}

// proxy implement
func NewFooPizzaStore() *PizzaStore {
	return &PizzaStore{Store: new(fooStore)}
}
