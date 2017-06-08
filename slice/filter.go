package slice

import (
	"errors"
	"reflect"
)

func Filter(src interface{}, filter func(i interface{}) bool) error {
	srcValueOf := reflect.ValueOf(src)
	if srcValueOf.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if srcValueOf.IsNil() {
		return errors.New("nil pointer passed to StructScan destination")
	}
	destRef := reflect.Indirect(srcValueOf)
	if destRef.Kind() != reflect.Slice {
		return errors.New("must pass a slice pointer with src")
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
