package implement

import (
	"github.com/LFZJun/go-lib/factory/logic"
)

// factory implement
type fooStore struct{}

func (ps *fooStore) CreatePizza() logic.Pizza {
	return new(FooPizza)
}

// proxy implement
func NewFooPizzaStore() *logic.PizzaStore {
	return &logic.PizzaStore{Store: new(fooStore)}
}
