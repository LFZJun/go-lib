package reflectl

import (
	"reflect"
)

// Deref is Indirect for reflect.Types
func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func IsSlicePtr(i interface{}) (reflect.Value, error) {
	srcValueOf := reflect.ValueOf(i)
	if srcValueOf.Kind() != reflect.Ptr {
		return reflect.Value{}, MustPtr
	}
	if srcValueOf.IsNil() {
		return reflect.Value{}, NilPtr
	}
	destRef := reflect.Indirect(srcValueOf)
	if destRef.Kind() != reflect.Slice {
		return reflect.Value{}, MustSlice
	}
	return destRef, nil
}

func TypeEqual(a, b reflect.Type) bool {
	switch {
	case a == nil, b == nil:
		return a == b
	case a == b:
		return true
	}
	switch a.Kind() {
	case reflect.Interface:
		return b.Implements(a)
	case reflect.Ptr:
		return a.AssignableTo(b) && a.ConvertibleTo(b)
	default:
		return a.Kind() == b.Kind()
	}
}

func TypeEqualI(a, b interface{}) bool {
	return TypeEqual(reflect.TypeOf(a), reflect.TypeOf(b))
}
