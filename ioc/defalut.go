package ioc

import (
	"log"
	"os"
)

var (
	logger           = log.New(os.Stdout, "[ioc] ", log.Lshortfile)
	defaultContainer = NewContainer()
	OpenPlugin       = true
)

func Add(water Water) {
	defaultContainer.Add(water)
}

func AddWithName(water Water, name string) {
	defaultContainer.AddWithName(water, name)
}

func GetWithName(name string) Water {
	return defaultContainer.GetWaterWithName(name)
}

func RegisterPlugin(lifecycle Lifecycle, p Plugin) {
	defaultContainer.RegisterPlugin(lifecycle, p)
}

func Start() {
	defaultContainer.Start()
}
