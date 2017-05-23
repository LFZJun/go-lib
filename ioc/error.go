package ioc

import "fmt"

type errorIoc int

const (
	ErrorPtr errorIoc = iota
	ErrorMissing
	ErrorUnexported
	ErrorTagDotIndex
	ErrorTagPtr
)

var errorString = [...]string{
	ErrorPtr:         "请传入Ptr \n当前类型 %v",
	ErrorMissing:     "找不到 %v",
	ErrorUnexported:  "%v: %v %v %v 需要大写变量首字母 mustBeExported",
	ErrorTagDotIndex: "错误field: %v %v `%v`  %v",
	ErrorTagPtr:      "错误field: %v %v `%v` 类型必须是Ptr",
}

func (e errorIoc) String() string {
	return errorString[e]
}

func (e errorIoc) Error() string {
	return errorString[e]
}

func (e errorIoc) Panic(argv ...interface{}) string {
	return fmt.Sprintf(e.String(), argv...)
}
