package reflectl

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(GetInterfaceDefaultName(1))
}
