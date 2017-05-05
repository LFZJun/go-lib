package parser

import (
	"fmt"
	"testing"
)

func TestCalculator_Evaluate(t *testing.T) {
	var expression Calculator = "1 + (1 + 2 * 1)"
	fmt.Println(expression.Evaluate())
}
