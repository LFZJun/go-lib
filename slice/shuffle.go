package slice

import (
	"math/rand"
	"time"
)

type Poker interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(poker Poker) {
	l := poker.Len()
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	for i := l; i > 0; {
		j := r.Intn(i)
		i--
		poker.Swap(i, j)
	}
}
