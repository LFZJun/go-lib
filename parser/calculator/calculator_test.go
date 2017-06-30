package calculator

import (
	"fmt"
	"testing"
	"strings"
)

func TestCalculator_Evaluate(t *testing.T) {
	var expression Calculator = "1 + (1 + 2 * 1)"
	fmt.Println(expression.Evaluate())
}

func TestLex(t *testing.T) {
	for token := range lexToml(strings.NewReader("1 + (1 + 2 * 1)")) {
		fmt.Println(token)
	}
}
