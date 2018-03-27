package topological

import (
	"fmt"
	"testing"
	"time"
)

func getOperatorMap() *OperatorMap {
	operatorMap := NewOperatorMap(10)

	action := func(o *OperatorNode) {
		fmt.Println(o.Val)
	}

	actionParallels := func(o *OperatorNode) {
		time.Sleep(time.Second)
		fmt.Println(o.Val)
	}

	root := &OperatorNode{
		Val:    "root",
		Action: action,
	}

	foo := &OperatorNode{
		Val:    "foo",
		Action: actionParallels,
	}
	bar := &OperatorNode{
		Val:    "bar",
		Action: actionParallels,
	}

	operatorMap.AddEdge(root, foo)
	operatorMap.AddEdge(root, bar)
	return operatorMap
}

func TestOperatorMap_Run(t *testing.T) {
	getOperatorMap().Run()
}
