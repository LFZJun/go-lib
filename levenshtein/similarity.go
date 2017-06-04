package levenshtein

import (
	"fmt"
	"math"
)

func Similarity(v1, v2 string) ln {
	distance := OptimizeDynamicDistance(v1, v2)
	return ln(1 - float64(distance)/math.Max(float64(len([]rune(v1))), float64(len([]rune(v2)))))
}

type ln float64

func (l ln) String() string {
	return fmt.Sprintf("%.2f%%", float64(l)*100)
}
