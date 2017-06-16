package reflectl

import "reflect"

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

func EqualType(a, b reflect.Type) bool {
	switch {
	case a == nil, b == nil:
		if a == b {
			return true
		}
	case a == b:
		return true
	}
	switch a.Kind() {
	case reflect.Interface:
		if b.Implements(a) {
			return true
		}
	case reflect.Ptr:
		if a.AssignableTo(b) && a.ConvertibleTo(b) {
			return true
		}
	default:
		if a.Kind() == b.Kind() {
			return true
		}
	}
	return false
}
