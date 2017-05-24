package main

import "fmt"

func Mod(shared, i int) int {
	fmt.Println(shared-1, i)
	fmt.Printf("%b %b\n", shared-1, i)
	return i & (shared - 1)
}

func main() {
	shared := 2 << 3
	i := 100
	fmt.Printf("%b\n", Mod(shared, i))
}
