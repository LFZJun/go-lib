package slice

import (
	"github.com/ljun20160606/go-lib/reflectl"
	"reflect"
)

func GroupBy(src interface{}, hash func(h interface{}) interface{}) error {
	destRef, err := reflectl.IsSlicePtr(src)
	if err != nil {
		return err
	}
	length := destRef.Len()
	set := make(map[interface{}]interface{}, length)
	tempSlice := reflect.MakeSlice(destRef.Type(), 0, length)
	for i := 0; i < length; i++ {
		v := destRef.Index(i).Interface()
		id := hash(v)
		if _, has := set[id]; !has {
			tempSlice = reflect.Append(tempSlice, reflect.ValueOf(v))
			set[id] = v
		}
	}
	destRef.Set(tempSlice)
	return nil
}
