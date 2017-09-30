package tpool

import (
	"testing"
	"fmt"
)

func TestSTATE(t *testing.T) {
	fmt.Printf("%b\n", CAPACITY)
	fmt.Printf("%b\n", RUNNING)
	fmt.Println(len(fmt.Sprintf("%b", RUNNING)))
	fmt.Printf("%b\n", ^RUNNING)
	fmt.Println(len(fmt.Sprintf("%b", ^RUNNING)))
	fmt.Printf("%b\n", SHUTDOWN)
	fmt.Printf("%b\n", STOP)
	fmt.Printf("%b\n", TIDYING)
	fmt.Printf("%b\n", TERMINATED)
	fmt.Printf("%b\n", ^CAPACITY)
	fmt.Printf("%b\n", runStateOf(RUNNING))
	fmt.Printf("%b\n", workerCountOf(RUNNING+2))
}
