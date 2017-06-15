package slice

import (
	"github.com/LFZJun/go-lib/reflectl"
	"reflect"
)

func GroupBy(src interface{}, hash func(h interface{}) interface{}, cmp func(i interface{}, j interface{}) bool) error {
	destRef, err := reflectl.IsSlicePtr(src)
	if err != nil {
		return err
	}
	set := make(map[interface{}]interface{})
	length := destRef.Len()
	for i := 0; i < length; i++ {
		v := destRef.Index(i).Interface()
		id := hash(v)
		if vv, has := set[id]; has {
			if cmp(v, vv) {
				set[id] = v
			}
		} else {
			set[id] = v
		}
	}
	tempSlice := reflect.MakeSlice(destRef.Type(), 0, 0)
	for _, v := range set {
		tempSlice = reflect.Append(tempSlice, reflect.ValueOf(v))
	}
	destRef.Set(tempSlice)
	return nil
}
