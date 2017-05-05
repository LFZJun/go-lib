package parser

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func errorType(i interface{}) {
	panic(fmt.Sprintf("类型错误 %v %v", reflect.TypeOf(i), i))
}

func newStack() *stack {
	return new(stack)
}

type stack struct {
	list []interface{}
	len  int
}

func (s *stack) Append(i interface{}) *stack {
	if i == nil {
		panic("不能存储nil")
	}
	s.list = append(s.list, i)
	s.len++
	return s
}

func (s *stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	back := s.list[s.len-1]
	s.list = s.list[:s.len-1]
	s.len--
	return back
}

func (s *stack) PopString() string {
	v := s.Pop()
	if v == nil {
		panic("无")
	}
	vv, ok := v.(string)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *stack) PopFloat64() float64 {
	v := s.Pop()
	vv, ok := v.(float64)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *stack) Back() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.list[s.len-1]
}

func (s *stack) BackString() string {
	v := s.Back()
	vv, ok := v.(string)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *stack) BackFloat64() float64 {
	v := s.Back()
	vv, ok := v.(float64)
	if !ok {
		errorType(v)
	}
	return vv
}

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

func applyOperation(operationStack *stack, outStack *stack) {
	outStack.Append(calc(outStack.PopFloat64(), outStack.PopFloat64(), operationStack.PopString()))
}

func isDigit(v int32) bool {
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
			case isDigit(v) || v == '.':
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
	operationStack := newStack()
	outStack := newStack()
	for token := range c.Parse() {
		switch {
		case numericValue.Match([]byte(token)):
			v, err := strconv.ParseFloat(token, 64)
			if err != nil {
				errorType(v)
			}
			outStack.Append(v)
		case token == "(":
			operationStack.Append(token)
		case token == ")":
			for operationStack.len > 0 && operationStack.Back() != "(" {
				applyOperation(operationStack, outStack)
			}
			operationStack.Pop()
		default:
			for operationStack.len > 0 && isOperator(operationStack.BackString()) && higherPriority(operationStack.BackString(), token) {
				applyOperation(operationStack, outStack)
			}
			operationStack.Append(token)
		}
	}
	for operationStack.len > 0 {
		applyOperation(operationStack, outStack)
	}
	return outStack.BackFloat64()
}
