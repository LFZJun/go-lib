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
	for token := range lex(strings.NewReader("1 - (1 + 2 * 1)")) {
		fmt.Println(token.String(), token.typ, token.Line, token.Col)
	}
}

func TestParser(t *testing.T) {
	fmt.Println(Parse(lex(strings.NewReader("1+11+1"))))
}
