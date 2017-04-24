package ioc

import (
	"reflect"
)

type field struct {
	structField reflect.StructField
	value       reflect.Value
	fieldOption *fieldOption
	parent      *holder
}

func (f *field) loadPlugin(p Plugin) {
	pValue := reflect.ValueOf(p.Value(f.fieldOption.path))
	f.value.Set(pValue)
}
