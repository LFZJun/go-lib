package ioc

import (
	"github.com/ljun20160606/go-lib/reflectl"
	"reflect"
	"strings"
)

type dropInfo struct {
	auto      bool
	dependent bool
	name      string
	prefix    string
	path      string
}

// 正则表达式说明
// 1. ioc:"[^\.]*" 由ioc控制依赖注入
// 2. ioc:"[^\.]*\..*" 由ioc plugin控制依赖注入
func newDropInfo(tagIoc string, class reflect.StructField) *dropInfo {
	si := &dropInfo{}
	dotIndex := strings.Index(tagIoc, ".")
	switch {
	case dotIndex == 0:
		panic(ErrorTagDotIndex.Panic(class.Name, class.Type, class.Tag, ". 不能放在首位"))
	case dotIndex == len(tagIoc)-1:
		panic(ErrorTagDotIndex.Panic(class.Name, class.Type, class.Tag, ". 不能放在末尾"))
	case dotIndex > 0: // 2.
		si.prefix = tagIoc[:dotIndex]
		si.path = tagIoc[dotIndex+1:]
		si.name = reflectl.GetTypeDefaultName(class.Type)
		return si
	}
	// 1.
	if tagIoc == "*" {
		si.auto = true
		si.dependent = true
		si.name = reflectl.GetTypeDefaultName(class.Type)
		return si
	}
	si.dependent = true
	si.name = tagIoc
	return si
}
