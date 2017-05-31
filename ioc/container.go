package ioc

import (
	"reflect"
	"sort"

	"github.com/LFZJun/go-lib/ioc/reflectl"
)

func NewContainer() Container {
	return &container{
		holderMap: make(map[string][]*Holder),
		fields:    make(map[string][]*field),
		plugins:   make(map[Lifecycle]plugins),
	}
}

type (
	Container interface {
		RegisterPlugin(lifecycle Lifecycle, p Plugin)
		Put(stone Stone)
		PutWithName(stone Stone, name string)
		GetHolder(name string, t reflect.Type) (h *Holder)
		GetStoneWithName(name string) Stone
		Start()
		putField(d *field)
	}

	container struct {
		holderMap map[string][]*Holder
		fields    map[string][]*field
		plugins   map[Lifecycle]plugins
	}
)

func (c *container) RegisterPlugin(lifecycle Lifecycle, p Plugin) {
	if _, ok := c.plugins[lifecycle]; !ok {
		c.plugins[lifecycle] = []Plugin{}
	}
	c.plugins[lifecycle] = append(c.plugins[lifecycle], p)
}

func (c *container) putStone(stone Stone, name string) {
	v := reflect.ValueOf(stone)
	t := v.Type()
	switch kind := t.Kind(); kind {
	case reflect.Ptr:
		if name == "" {
			name = reflectl.GetValueDefaultName(v)
		}
	default:
		panic(ErrorPtr.Panic(kind))
	}
	logger.Printf("放入 %v", name)
	// 额，没想到并发的场景所以没加锁
	if _, ok := c.holderMap[name]; !ok {
		holder := newHolder(stone, t, v, c)
		c.holderMap[name] = append(c.holderMap[name], holder)
	}
}

func (c *container) Put(stone Stone) {
	c.putStone(stone, "")
}

func (c *container) PutWithName(stone Stone, name string) {
	c.putStone(stone, name)
}

func (c *container) GetStoneWithName(name string) Stone {
	if hs, ok := c.holderMap[name]; ok {
		switch len(hs) {
		case 0:
			return nil
		default:
			return hs[0].Stone
		}
	}
	return nil
}

func (c *container) GetHolder(name string, t reflect.Type) (h *Holder) {
	if holder, found := c.holderMap[name]; found {
		for _, h := range holder {
			if h.Equal(t) {
				return h
			}
		}
	}
	return nil
}

func (c *container) putField(d *field) {
	prefix := d.fieldOption.prefix
	if _, has := c.fields[prefix]; !has {
		c.fields[prefix] = []*field{}
	}
	c.fields[prefix] = append(c.fields[prefix], d)
}

func (c *container) eachHolder(holderFunc HolderFunc) {
	for _, v := range c.holderMap {
		for _, vv := range v {
			holderFunc(vv)
		}
	}
}

func (c *container) genDependents() {
	c.eachHolder(func(holder *Holder) {
		holder.genDependents()
	})
}

func (c *container) loadPlugins(lifecycle Lifecycle) {
	ps, ok := c.plugins[lifecycle]
	if !ok || len(ps) == 0 {
		return
	}
	sort.Sort(ps)
	for _, p := range ps {
		fs, ok := c.fields[p.Prefix()]
		if !ok {
			return
		}
		for _, f := range fs {
			f.loadPlugin(p)
		}
	}
}

func (c *container) Start() {
	c.genDependents()
	c.loadPlugins(BeforeInit)
	c.init()
	c.loadPlugins(BeforeReady)
	c.ready()
}

func (c *container) init() {
	set := make(HolderSet)
	c.eachHolder(func(holder *Holder) {
		holder.init(set)
	})
}

func (c *container) ready() {
	set := make(HolderSet)
	c.eachHolder(func(holder *Holder) {
		holder.ready(set)
	})
}
