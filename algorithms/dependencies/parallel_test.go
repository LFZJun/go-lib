package dependencies

import (
	"fmt"
	"testing"
)

func getOperatorMap() *OperatorMap {
	operatorMap := NewOperatorMap(10)

	action := func(o *Operator) {
		fmt.Println(o.ID)
	}

	foo := &Operator{
		ID:     "foo",
		Action: action,
	}
	bar := &Operator{
		ID:     "bar",
		Action: action,
	}

	operatorMap.AddEdge(foo, bar)
	return operatorMap
}

func TestOperatorMap_Run(t *testing.T) {
	getOperatorMap().Run()
}
