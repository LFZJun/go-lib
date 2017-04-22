package ioc

import (
	"fmt"
	"reflect"

)

type Holder struct {
	Stone      Stone
	Class      reflect.Type // class.kind() Ptr
	Value      reflect.Value // value.kind() Ptr
	Parent     *Container
	Dependents []*Holder
}

func newHolder(stone Stone, typee reflect.Type, value reflect.Value, container *Container) *Holder {
	return &Holder{
		Stone:      stone,
		Class:      typee,
		Value:      value,
		Parent:     container,
		Dependents: []*Holder{},
	}
}

func (h *Holder) genDependents() {
	classElem := h.Class.Elem()
	valueElem := h.Value.Elem()
	for num := valueElem.NumField() - 1; num >= 0; num-- {
		class := classElem.Field(num)
		value := valueElem.Field(num)
		tag, ok := class.Tag.Lookup("ioc")
		if !ok || tag == "-" {
			continue
		}
		if !value.CanSet() {
			panic(fmt.Sprintf("不能修改 %v", class))
		}
		field := buildFieldOptions(tag, class)
		if !field.dependent {
			h.Parent.putDelayFields(&delayField{
				class:       class,
				value:       value,
				fieldOption: field,
				parent:      h,
			})
			return
		}
		holder := h.Parent.GetHolder(field.name, value.Type())
		if holder == nil {
			panic(fmt.Sprintf("找不到 %v", value.Type()))
		}
		h.Dependents = append(h.Dependents, holder)
		value.Set(holder.Value)
	}
}
