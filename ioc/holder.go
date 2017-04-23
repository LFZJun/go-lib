package ioc

import (
	"fmt"
	"reflect"
)

type Holder struct {
	Stone      Stone
	Class      reflect.Type  // class.kind() Ptr
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

func (h *Holder) equal(t reflect.Type) Stone {
	switch t.Kind() {
	case reflect.Interface:
		if h.Class.Implements(t) {
			return h.Stone
		}
	case reflect.Struct:
		t = reflect.PtrTo(t)
		fallthrough
	case reflect.Ptr:
		if h.Class.AssignableTo(t) && h.Class.ConvertibleTo(t) {
			return h.Stone
		}
	}
	return nil
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

func (h *Holder) init(set HolderSet) {
	if _, has := set[h]; has {
		return
	}
	set[h] = struct{}{}
	for _, dependence := range h.Dependents {
		dependence.init(set)
	}
	if init, ok := h.Stone.(Init); ok {
		init.Init()
	}
}

func (h *Holder) ready(set HolderSet) {
	if _, has := set[h]; has {
		return
	}
	set[h] = struct{}{}
	for _, dependence := range h.Dependents {
		dependence.ready(set)
	}
	if init, ok := h.Stone.(Ready); ok {
		init.Ready()
	}
}
