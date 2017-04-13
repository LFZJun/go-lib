package main

import (
	"fmt"
	"github.com/LFZJun/go-lib/factory/implement"
)

func main() {
	fooPizzaStore := implement.NewFooPizzaStore()
	fooPizzaStore.OrderPizza()
	fmt.Println(fooPizzaStore)
}
