package ioc

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[ioc]", log.LstdFlags)
var defaultContainer = NewContainer()

func Put(stone Stone) {
	defaultContainer.Put(stone)
}

func PutWithName(stone Stone, name string) {
	defaultContainer.PutWithName(stone, name)
}

func GetStoneWithName(name string) Stone {
	return defaultContainer.GetStoneWithName(name)
}

func RegisterPlugin(lifecycle Lifecycle, p Plugin) {
	defaultContainer.registerPlugin(lifecycle, p)
}

func Start() {
	defaultContainer.Start()
}
