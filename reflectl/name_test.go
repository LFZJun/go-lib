package reflectl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assert.Equal(t, "int", GetInterfaceDefaultName(1))
}
