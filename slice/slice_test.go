package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type foo struct {
	Id    int
	Value int64
}

func TestGroupBy(t *testing.T) {
	srcFoo := []foo{{1, 1}, {1, 3}, {1, 4}}
	hashFunc := func(h interface{}) interface{} {
		return h.(foo).Id
	}
	err := GroupBy(&srcFoo, hashFunc)
	assert.NoError(t, err)
	assert.Equal(t, []foo{{1, 1}}, srcFoo)
}

func TestFilter(t *testing.T) {
	srcFoo := []foo{{1, 1}, {1, 3}, {1, 4}}
	filterFunc := func(i interface{}) bool {
		t := i.(foo)
		return t.Value < 4
	}
	err := Filter(&srcFoo, filterFunc)
	assert.NoError(t, err)
	assert.Equal(t, []foo{{1, 1}, {1, 3}}, srcFoo)
}

type poker []int

func (p poker) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p poker) Len() int {
	return len(p)
}

func TestShuffle(t *testing.T) {
	var foo poker = []int{1, 2, 3, 4}
	assert.NotPanics(t, func() {
		Shuffle(foo)
	})
}
