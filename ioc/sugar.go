package ioc

import (
	"reflect"
)

type (
	Sugar interface {
		Container() Container
		ParentCup() *Cup
		Value() reflect.Value
		Type() reflect.Type
		Water() Water
		LoadPlugin(p Plugin)
		Prefix() string
	}

	sugar struct {
		classInfo reflect.StructField
		parentCup *Cup
		value     reflect.Value
		dropInfo  *dropInfo
	}
)

func (s *sugar) LoadPlugin(p Plugin) {
	s.value.Set(reflect.ValueOf(p.Lookup(s.dropInfo.path, s)))
}

func (s *sugar) Container() Container {
	return s.parentCup.Container
}

func (s *sugar) ParentCup() *Cup {
	return s.parentCup
}

func (s *sugar) Value() reflect.Value {
	return s.value
}

func (s *sugar) Type() reflect.Type {
	return s.value.Type()
}

func (s *sugar) Water() Water {
	return s.value.Interface()
}

func (s *sugar) Prefix() string {
	return s.dropInfo.prefix
}
