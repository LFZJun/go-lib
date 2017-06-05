package slice

import (
	"errors"
	"reflect"
)

func GroupBy(dest, src interface{}, hash func(h interface{}) interface{}, cmp func(i interface{}, j interface{}) bool) error {
	destValueOf := reflect.ValueOf(dest)
	if destValueOf.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if destValueOf.IsNil() {
		return errors.New("nil pointer passed to StructScan destination")
	}
	destRef := reflect.Indirect(destValueOf)
	if destRef.Kind() != reflect.Slice {
		return errors.New("must pass a slice pointer with dest")
	}
	srcRef := reflect.ValueOf(src)
	if destRef.Type() != srcRef.Type() {
		return errors.New("must pass same type of dest, src")
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
