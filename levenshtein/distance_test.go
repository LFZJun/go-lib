package levenshtein

import (
	"fmt"
	"testing"
)

func TestSimilarity(t *testing.T) {
	fmt.Println(Similarity("acc", "adc"))
}

func TestDistance(t *testing.T) {
	fmt.Println(Distance("acc", "abc"))
}

func TestDistanceDivide(t *testing.T) {
	fmt.Println(DistanceDivide("acc", "abc"))
}
