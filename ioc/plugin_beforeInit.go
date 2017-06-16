package ioc

import (
	"reflect"
	"github.com/LFZJun/go-lib/reflectl"
)

func init() {
	if OpenPlugin {
		LoadBeforeInitPlugin()
	}
}

func LoadBeforeInitPlugin() {
	RegisterPlugin(BeforeInit, new(BeforeInitPlugin))
}

var (
	beforeInitType = reflect.TypeOf((*BeforeInitType)(nil)).Elem()
)

type (
	BeforeInitType interface {
		BeforeInit() interface{}
	}
	BeforeInitPlugin struct {
	}
)

func (b *BeforeInitPlugin) Lookup(path string, sugar Sugar) (v interface{}) {
	if path == "*" {
		sugar.Container().EachCup(func(cup *Cup) bool {
			if beforeInitType, ok := cup.Water.(BeforeInitType); ok {
				vv := beforeInitType.BeforeInit()
				if reflectl.EqualType(sugar.Type(), reflect.TypeOf(vv)) {
					v = vv
					return true
				}
			}
			return false
		})
		if v == nil {
			panic(ErrorMissing.Panic(b.Prefix() + "." + path))
		}
		return v
	}
	cup := sugar.Container().GetCup(path, beforeInitType)
	if cup == nil {
		panic(ErrorMissing.Panic(b.Prefix() + "." + path))
	}
	return cup.Water.(BeforeInitType).BeforeInit()
}

func (b *BeforeInitPlugin) Prefix() string {
	return "@"
}

func (b *BeforeInitPlugin) Priority() int {
	return 1
}
