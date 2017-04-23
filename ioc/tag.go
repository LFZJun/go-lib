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

type delayField struct {
	class       reflect.StructField
	value       reflect.Value
	fieldOption *fieldOption
	parent      *Holder
}

func buildFieldOptions(tag string, class reflect.StructField) *fieldOption {
	to := &fieldOption{}
	dotIndex := strings.Index(tag, ".")
	switch {
	case dotIndex > 0:
		to.prefix = tag[:dotIndex]
		to.path = tag[dotIndex+1:]
		return to
	case dotIndex == 0:
		panic(fmt.Sprintf("错误field: %v %v `%v`  . 不能放在首位", class.Name, class.Type, class.Tag))
	}
	if class.Type.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("错误field: %v %v `%v` 类型必须是Ptr", class.Name, class.Type, class.Tag))
	}
	if tag == "*" {
		to.auto = true
		to.dependent = true
		to.name = class.Type.Elem().Name()
		return to
	}
	to.dependent = true
	to.name = tag
	return to
}
