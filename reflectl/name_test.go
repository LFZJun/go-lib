package reflectl

import (
	"fmt"
	"github.com/cocotyty/summer"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(GetInterfaceDefaultName(1))
}
