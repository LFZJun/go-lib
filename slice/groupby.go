package slice

import (
	"reflect"
)

func GroupBy(dest, src interface{}, hash func(h interface{}) interface{}, cmp func(i interface{}, j interface{}) bool) error {
	destValueOf := reflect.ValueOf(dest)
	if destValueOf.Kind() != reflect.Ptr {
		return MustPtr
	}
	if destValueOf.IsNil() {
		return NilPtr
	}
	destRef := reflect.Indirect(destValueOf)
	if destRef.Kind() != reflect.Slice {
		return MustSlice
	}
	srcRef := reflect.ValueOf(src)
	if destRef.Type() != srcRef.Type() {
		return MustSameType
	}
	set := make(map[interface{}]interface{})
	length := srcRef.Len()
	for i := 0; i < length; i++ {
		v := srcRef.Index(i).Interface()
		id := hash(v)
		if vv, has := set[id]; has {
			if cmp(v, vv) {
				set[id] = v
			}
		} else {
			set[id] = v
		}
	}
	for _, v := range set {
		destRef.Set(reflect.Append(destRef, reflect.ValueOf(v)))
	}
	return nil
}
