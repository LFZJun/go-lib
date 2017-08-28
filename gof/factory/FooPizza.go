package factory

import (
	"fmt"
)

// worker implement
type FooPizza struct{}

func (fp *FooPizza) Prepare() {
	fmt.Println("foo pizza prepare")
}

func (fp *FooPizza) Bake() {
	fmt.Println("foo pizza bake")
}

func (fp *FooPizza) Cut() {
	fmt.Println("foo pizza cut")
}

func (fp *FooPizza) Box() {
	fmt.Println("foo pizza box")
}
