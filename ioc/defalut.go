package ioc

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[ioc]", log.LstdFlags)
var container = NewContainer()

func Put(stone Stone) {
	container.Put(stone)
}

func Start()  {
	container.Start()
}
