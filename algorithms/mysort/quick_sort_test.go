package mysort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSortV1(t *testing.T) {
	assert.Equal(t, []int{1, 4, 6, 7}, quickSortV1([]int{6, 1, 7, 4}))
}

func TestQuickSortV2(t *testing.T) {
	assert.Equal(t, []int{1, 6, 7, 9, 4}, quickSortV2([]int{6, 9, 1, 7, 4}, 0, 3))
}

func TestQuickSort(t *testing.T) {
	arrays := []int{6, 1, 7, 4}
	quickSort(arrays)
	assert.Equal(t, []int{1, 4, 6, 7}, arrays)
}
