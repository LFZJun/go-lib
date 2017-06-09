package slice

import (
	"reflect"
)

func Filter(src interface{}, filter func(i interface{}) bool) error {
	srcValueOf := reflect.ValueOf(src)
	if srcValueOf.Kind() != reflect.Ptr {
		return MustPtr
	}
	if srcValueOf.IsNil() {
		return NilPtr
	}
	destRef := reflect.Indirect(srcValueOf)
	if destRef.Kind() != reflect.Slice {
		return MustSlice
	}
	tempSlice := reflect.MakeSlice(destRef.Type(), 0, 0)
	length := destRef.Len()
	for i := 0; i < length; i++ {
		v := destRef.Index(i)
		vv := v.Interface()
		if !filter(vv) {
			tempSlice = reflect.Append(tempSlice, v)
		}
	}
	destRef.Set(tempSlice)
	return nil
}
