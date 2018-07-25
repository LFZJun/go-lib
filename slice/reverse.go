package slice

import (
	"github.com/ljun20160606/go-lib/reflectl"
	"reflect"
)

func Reverse(s interface{}) error {
	valueOf := reflect.ValueOf(s)
	indirect := reflect.Indirect(valueOf)
	if indirect.Kind() != reflect.Slice {
		return reflectl.MustSlice
	}
	for i, j := 0, valueOf.Len()-1; i < j; i, j = i+1, j-1 {
		left := valueOf.Index(i)
		right := valueOf.Index(j)
		temp := reflect.ValueOf(left.Interface())
		left.Set(right)
		right.Set(temp)
	}
	return nil
}
