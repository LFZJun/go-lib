package reflectl

import "reflect"

func GetValueDefaultName(v reflect.Value) string {
	return reflect.Indirect(v).Type().String()
}

func GetInterfaceDefaultName(i interface{}) string {
	return GetValueDefaultName(reflect.ValueOf(i))
}

func GetTypeDefaultName(t reflect.Type) string {
	return Deref(t).String()
}
