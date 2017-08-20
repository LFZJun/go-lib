package stack

import (
	"fmt"
	"reflect"
)

func errorType(i interface{}) {
	panic(fmt.Sprintf("类型错误 %v %v", reflect.TypeOf(i), i))
}

// 链表的话 list.New()
func NewSimpleStack() *SimpleStack {
	return new(SimpleStack)
}

type SimpleStack struct {
	list []interface{}
	Len  int
}

func (s *SimpleStack) Append(i interface{}) *SimpleStack {
	if i == nil {
		panic("不能存储nil")
	}
	s.list = append(s.list, i)
	s.Len++
	return s
}

func (s *SimpleStack) Pop() interface{} {
	if s.Len == 0 {
		panic("Len 0")
	}
	back := s.list[s.Len-1]
	s.list = s.list[:s.Len-1]
	s.Len--
	return back
}

func (s *SimpleStack) PopByte() byte {
	v := s.Pop()
	vv, ok := v.(byte)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *SimpleStack) PopString() string {
	v := s.Pop()
	vv, ok := v.(string)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *SimpleStack) PopFloat64() float64 {
	v := s.Pop()
	vv, ok := v.(float64)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *SimpleStack) Back() interface{} {
	if s.Len == 0 {
		panic("Len 0")
	}
	return s.list[s.Len-1]
}

func (s *SimpleStack) BackString() string {
	v := s.Back()
	vv, ok := v.(string)
	if !ok {
		errorType(v)
	}
	return vv
}

func (s *SimpleStack) BackFloat64() float64 {
	v := s.Back()
	vv, ok := v.(float64)
	if !ok {
		errorType(v)
	}
	return vv
}
