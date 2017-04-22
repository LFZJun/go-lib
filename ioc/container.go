package ioc

import (
	"errors"
	"fmt"
	"reflect"
)

type Container struct {
	holderMap   map[string][]*Holder
	delayFields map[string][]*delayField
}

func NewContainer() *Container {
	return &Container{
		holderMap:   make(map[string][]*Holder),
		delayFields: make(map[string][]*delayField),
	}
}

func (c *Container) Put(stone Stone) {
	v := reflect.ValueOf(stone)
	t := v.Type()
	var name string
	switch kind := t.Kind(); kind {
	case reflect.Ptr:
		name = t.Elem().Name()
	default:
		panic(errors.New(fmt.Sprintf("请传入Ptr \n当前类型 %v", kind)))
	}
	//name = strings.ToLower(name[:1]) + name[1:]
	logger.Printf("放入 %v", name)
	if _, ok := c.holderMap[name]; !ok {
		holder := newHolder(stone, t, v, c)
		c.holderMap[name] = append(c.holderMap[name], holder)
	}
}

func (c *Container) GetHolder(name string, t reflect.Type) (h *Holder) {
	if holder, found := c.holderMap[name]; found {
		for _, h := range holder {
			if stone := c.findStone(h, t); stone != nil {
				return h
			}
		}
	}
	return nil
}

func (c *Container) findStone(h *Holder, t reflect.Type) Stone {
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

func (c *Container) putDelayFields(d *delayField) {
	prefix := d.fieldOption.prefix
	if _, has := c.delayFields[prefix]; !has {
		c.delayFields[prefix] = []*delayField{}
	}
	c.delayFields[prefix] = append(c.delayFields[prefix], d)
}

func (c *Container) eachHolder(funcc func(holder *Holder)) {
	for _, v := range c.holderMap {
		for _, vv := range v {
			funcc(vv)
		}
	}
}

func (c *Container) genDependents() {
	c.eachHolder(func(holder *Holder) {
		holder.genDependents()
	})
}

func (c *Container) Start() {
	c.genDependents()
}
