package ioc

import (
	"reflect"
	"sort"

	"fmt"
	"github.com/LFZJun/go-lib/reflectl"
)

func NewContainer() Container {
	return &container{
		cupMap:   make(map[string][]*Cup),
		sugarMap: make(map[string][]Sugar),
		plugins:  make(map[Lifecycle]Plugins),
	}
}

type (
	Container interface {
		RegisterPlugin(lifecycle Lifecycle, p Plugin)
		addSugar(sugar Sugar)
		Add(water Water)
		AddWithName(water Water, name string)
		GetCup(name string, dropType reflect.Type) (h *Cup)
		GetWaterWithName(name string) Water
		EachCup(cupFunc CupFunc)
		Start()
	}

	containPlugin interface {
	}
	container struct {
		cupMap   map[string][]*Cup
		sugarMap map[string][]Sugar
		plugins  map[Lifecycle]Plugins
	}
)

func (c *container) RegisterPlugin(lifecycle Lifecycle, p Plugin) {
	if _, ok := c.plugins[lifecycle]; !ok {
		c.plugins[lifecycle] = []Plugin{}
	}
	c.plugins[lifecycle] = append(c.plugins[lifecycle], p)
}

func (c *container) addSugar(sugar Sugar) {
	prefix := sugar.Prefix()
	if _, has := c.sugarMap[prefix]; !has {
		c.sugarMap[prefix] = []Sugar{sugar}
		return
	}
	c.sugarMap[prefix] = append(c.sugarMap[prefix], sugar)
}

func (c *container) addWater(water Water, name string) {
	v := reflect.ValueOf(water)
	t := v.Type()
	switch kind := t.Kind(); kind {
	case reflect.Ptr:
		if name == "" {
			name = reflectl.GetValueDefaultName(v)
		}
	default:
		panic(ErrorPtr.Panic(kind))
	}
	logger.Output(4, fmt.Sprintf("放入 %v", name))
	// 额，没想到并发的场景所以没加锁
	if _, ok := c.cupMap[name]; !ok {
		cup := newCup(water, t, v, c)
		c.cupMap[name] = append(c.cupMap[name], cup)
	}
}

func (c *container) Add(water Water) {
	c.addWater(water, "")
}

func (c *container) AddWithName(water Water, name string) {
	c.addWater(water, name)
}

func (c *container) GetWaterWithName(name string) Water {
	if hs, ok := c.cupMap[name]; ok {
		switch len(hs) {
		case 0:
			return nil
		default:
			return hs[0].Water
		}
	}
	return nil
}

func (c *container) GetCup(name string, t reflect.Type) (h *Cup) {
	if cups, found := c.cupMap[name]; found {
		for _, cup := range cups {
			if reflectl.TypeEqual(t, cup.Class) {
				return cup
			}
		}
	}
	return nil
}

func (c *container) EachCup(cupFunc CupFunc) {
	for _, v := range c.cupMap {
		for _, vv := range v {
			if cupFunc(vv) {
				return
			}
		}
	}
}

func (c *container) injectDependency() {
	c.EachCup(func(cup *Cup) bool {
		cup.injectDependency()
		return false
	})
}

func (c *container) loadPlugins(lifecycle Lifecycle) {
	plugins, ok := c.plugins[lifecycle]
	if !ok || len(plugins) == 0 {
		return
	}
	sort.Sort(plugins)
	for _, plugin := range plugins {
		sugars, ok := c.sugarMap[plugin.Prefix()]
		if !ok {
			continue
		}
		for _, sugar := range sugars {
			sugar.LoadPlugin(plugin)
		}
	}
}

func (c *container) Start() {
	c.injectDependency()
	c.loadPlugins(BeforeInit)
	c.init()
	c.loadPlugins(BeforeReady)
	c.ready()
}

func (c *container) init() {
	set := make(CupSet)
	c.EachCup(func(cup *Cup) bool {
		cup.init(set)
		return false
	})
}

func (c *container) ready() {
	set := make(CupSet)
	c.EachCup(func(cup *Cup) bool {
		cup.ready(set)
		return false
	})
}
