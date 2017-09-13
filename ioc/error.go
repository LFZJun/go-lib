package ioc

import "fmt"

type errorIoc int

const (
	ErrorType errorIoc = iota
	ErrorMissing
	ErrorUnexported
	ErrorTagDotIndex
	ErrorStopIterator
)

var errorString = [...]string{
	ErrorType:         "请传入Ptr \n当前basic type %v",
	ErrorMissing:      "找不到 %v",
	ErrorUnexported:   "%v: %v %v %v 需要大写变量首字母 mustBeExported",
	ErrorTagDotIndex:  "错误field: %v %v `%v`  %v",
	ErrorStopIterator: "停止循环",
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
