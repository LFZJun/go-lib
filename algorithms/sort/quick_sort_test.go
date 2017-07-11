package sort

import (
	"testing"
	"fmt"
)

func TestQuickSortV1(t *testing.T) {
	fmt.Println(quickSortV1([]int{6, 1, 7, 4}))
}

func TestQuickSortV2(t *testing.T) {
	fmt.Println(quickSortV2([]int{6, 9, 1, 7, 4}, 0, 3))
}
