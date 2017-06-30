package calculator

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/LFZJun/go-lib/algorithms/stack"
)

const (
	OPERATION = "+-/*"
)

var (
	numericValue      = regexp.MustCompile(`\d+(\.\d+)?`)
	operationPriority = map[string]int{
		"+": 0,
		"-": 0,
		"*": 1,
		"/": 1,
	}
)

func isOperator(tokenValue string) bool {
	return strings.Contains(OPERATION, tokenValue)
}

func higherPriority(op1 string, op2 string) bool {
	return operationPriority[op1] >= operationPriority[op2]
}

func calc(n2 float64, n1 float64, operator string) float64 {
	switch operator {
	case "-":
		return n1 - n2
	case "+":
		return n1 + n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	}
	panic(fmt.Sprintf("无法识别运算符 %v", operator))
}

func applyOperation(operationStack *stack.SimpleStack, outStack *stack.SimpleStack) {
	outStack.Append(calc(outStack.PopFloat64(), outStack.PopFloat64(), operationStack.PopString()))
}

func isDigitt(v int32) bool {
	if v >= '0' && v <= '9' {
		return true
	}
	return false
}

func isBlank(v byte) bool {
	if v == ' ' || v == '\t' {
		return true
	}
	return false
}

type Calculator string

func (c Calculator) Parse() <-chan string {
	s := make(chan string)
	go func() {
		number := bytes.NewBuffer(nil)
		for _, v := range string(c) {
			switch {
			case isDigitt(v) || v == '.':
				number.WriteByte(byte(v))
			default:
				if number.Len() > 0 {
					s <- number.String()
					number.Reset()
				}
				if !isBlank(byte(v)) {
					s <- string(v)

				}
			}

		}
		if number.Len() > 0 {
			s <- number.String()
		}
		close(s)
	}()
	return s
}

func (c Calculator) Evaluate() float64 {
	operationStack := stack.NewStack()
	outStack := stack.NewStack()
	for token := range c.Parse() {
		switch {
		case numericValue.Match([]byte(token)):
			v, err := strconv.ParseFloat(token, 64)
			if err != nil {
				panic(v)
			}
			outStack.Append(v)
		case token == "(":
			operationStack.Append(token)
		case token == ")":
			for operationStack.Len > 0 && operationStack.Back() != "(" {
				applyOperation(operationStack, outStack)
			}
			operationStack.Pop()
		default:
			for operationStack.Len > 0 && isOperator(operationStack.BackString()) && higherPriority(operationStack.BackString(), token) {
				applyOperation(operationStack, outStack)
			}
			operationStack.Append(token)
		}
	}
	for operationStack.Len > 0 {
		applyOperation(operationStack, outStack)
	}
	return outStack.BackFloat64()
}
