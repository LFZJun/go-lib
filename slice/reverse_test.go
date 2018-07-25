package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	ast := assert.New(t)
	input := []interface{}{1, "2", 3, '4'}
	Reverse(input)
	ast.Equal([]interface{}{'4', 3, "2", 1}, input)
}
