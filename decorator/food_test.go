package decorator

import (
	"fmt"
	"testing"
)

func TestFood(t *testing.T) {
	ramen := new(Ramen)
	egg := Egg{Noddles: ramen}
	sausage := Sausage{Noddles: egg}

	fmt.Println(sausage.Description())
	fmt.Println(sausage.Price())
}
