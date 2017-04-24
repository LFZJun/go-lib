package ioc

import (
	"fmt"
	"reflect"
	"strings"
)

type fieldOption struct {
	auto      bool
	dependent bool
	name      string
	prefix    string
	path      string
}

func buildFieldOptions(tagIoc string, class reflect.StructField) *fieldOption {
	to := &fieldOption{}
	dotIndex := strings.Index(tagIoc, ".")
	switch {
	case dotIndex == 0:
		panic(fmt.Sprintf("错误field: %v %v `%v`  . 不能放在首位", class.Name, class.Type, class.Tag))
	case dotIndex == len(tagIoc) - 1:
		panic(fmt.Sprintf("错误field: %v %v `%v`  . 不能放在末尾", class.Name, class.Type, class.Tag))
	case dotIndex > 0:
		to.prefix = tagIoc[:dotIndex]
		to.path = tagIoc[dotIndex+1:]
		return to
	}
	if class.Type.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("错误field: %v %v `%v` 类型必须是Ptr", class.Name, class.Type, class.Tag))
	}
	if tagIoc == "*" {
		to.auto = true
		to.dependent = true
		to.name = class.Type.Elem().Name()
		return to
	}
	to.dependent = true
	to.name = tagIoc
	return to
}
