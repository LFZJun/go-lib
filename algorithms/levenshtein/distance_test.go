package levenshtein

import (
	"fmt"
	"testing"
)

func TestSimilarity(t *testing.T) {
	fmt.Println(Similarity("acc", "adc"))
}

// 分治法
func TestDivideDistance(t *testing.T) {
	fmt.Println(DivideDistance("acc", "abc"))
}

// 动态规划
func TestDynamicDistance(t *testing.T) {
	fmt.Println(DynamicDistance("acc", "abc"))
}

// 优化空间复杂度的动态规划
func TestDistance(t *testing.T) {
	fmt.Println(OptimizeDynamicDistance("acc", "abc"))
}
