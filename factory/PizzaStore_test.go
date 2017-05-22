package factory

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fooPizzaStore := NewFooPizzaStore()
	fooPizzaStore.OrderPizza()
	fmt.Println(fooPizzaStore)
}
