package parser

import (
	"strconv"
)

type StructTag string

func (t StructTag) Parse() map[string]string {
	table := make(map[string]string)

	length := len(t)
	i := 0
	index := 0
	isEOF := func() bool {
		return t[i:] == ""
	}

	errorr := func(position int, message string) {
		panic("位置" + strconv.Itoa(position) + ": " + string(t[:position]) + "^" + string(t[position:]) + " " + message)
	}

	next := func(b byte) {
		if i >= length || t[i] != b {
			errorr(i, "需要 "+string(b))
		}
	}

	white := func() {
		for i < len(t) && t[i] == ' ' {
			i++
		}
	}

	stringg := func() string {
		index = i
		for i < length && t[i] != '"' {
			if t[i] == '\\' {
				i++
			}
			i++
		}
		return string(t[index:i])
	}

	value := func(name string) {
		next('"')
		i++
		value := stringg()
		next('"')
		table[name] = string(value)
		i++
	}

	tag := func() {
		index = i
		for i < length && t[i] > ' ' && t[i] != ':' && t[i] != '"' && t[i] != 0x7f {
			i++
		}
		if index == i {
			errorr(index, " error: tag key不能为空")
		}
		name := string(t[index:i])
		if _, ok := table[name]; ok {
			errorr(index, " 重复key: "+name)
		}
		white()
		next(':')
		i++
		white()
		value(name)
	}

	white()
	for !isEOF() {
		white()
		tag()
	}
	return table
}
