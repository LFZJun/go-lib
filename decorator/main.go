package main

import (
	"fmt"
	. "ljun/go-lib/decorator/implement"
)

func main() {
	ramen := new(Ramen)
	egg := Egg{Noddles: ramen}
	sausage := Sausage{Noddles: egg}

	fmt.Println(sausage.Description())
	fmt.Println(sausage.Price())
}
