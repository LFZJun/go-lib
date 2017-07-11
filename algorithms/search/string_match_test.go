package search

import (
	"fmt"
	"testing"
)

func TestIndexOf(t *testing.T) {
	const tt = "abc"
	i := IndexOf(tt, "a")
	if i != nil {
		fmt.Println(tt[i[0]:i[1]])
	}
}
