package implement

import (
	"fmt"
)

type NoCommand struct {
}

func (n *NoCommand) Execute() {
	fmt.Println("nocommand execute")
}

func (n *NoCommand) Undo() {
	fmt.Println("nocommand undo")
}

type LightCommand struct {
	status bool
	Light  *Light
}

func (lc *LightCommand) Execute() {
	lc.status = !lc.status
	switch lc.status {
	case true:
		lc.Light.On()
	case false:
		lc.Light.Off()
	}
}

func (lc *LightCommand) Undo() {
	lc.Execute()
}
