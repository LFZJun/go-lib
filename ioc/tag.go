package ioc

import (
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
		panic(ErrorTagDotIndex.Panic(class.Name, class.Type, class.Tag, ". 不能放在首位"))
	case dotIndex == len(tagIoc)-1:
		panic(ErrorTagDotIndex.Panic(class.Name, class.Type, class.Tag, ". 不能放在末尾"))
	case dotIndex > 0:
		to.prefix = tagIoc[:dotIndex]
		to.path = tagIoc[dotIndex+1:]
		return to
	}
	if class.Type.Kind() != reflect.Ptr {
		panic(ErrorTagPtr.Panic(class.Name, class.Type, class.Tag))
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
