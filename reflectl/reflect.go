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
