package slice

import (
	"errors"
	"reflect"
)

func GroupBy(dest, src interface{}, hash func(h interface{}) interface{}, key func(k interface{}) int64, reverse bool) error {
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
	var compareing func(a, b int64) bool
	if reverse {
		compareing = func(a, b int64) bool {
			return a <= b
		}
	} else {
		compareing = func(a, b int64) bool {
			return a > b
		}
	}
	for i := 0; i < length; i++ {
		v := srcRef.Index(i).Interface()
		id := hash(v)
		t := key(v)
		if vv, has := set[id]; has {
			if compareing(key(vv), t) {
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
