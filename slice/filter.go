package slice

import (
	"github.com/ljun20160606/go-lib/reflectl"
	"reflect"
)

func Filter(src interface{}, filter func(i interface{}) bool) error {
	destRef, err := reflectl.IsSlicePtr(src)
	if err != nil {
		return err
	}
	tempSlice := reflect.MakeSlice(destRef.Type(), 0, 0)
	length := destRef.Len()
	for i := 0; i < length; i++ {
		v := destRef.Index(i)
		vv := v.Interface()
		if filter(vv) {
			tempSlice = reflect.Append(tempSlice, v)
		}
	}
	destRef.Set(tempSlice)
	return nil
}
