package main

import (
	"github.com/LFZJun/go-lib/command/implement"
	"github.com/LFZJun/go-lib/command/logic"
)

func main() {
	light := &implement.Light{Site: "hall"}
	lightCommand := &implement.LightCommand{Light: light}
	party := &implement.Party{Commands: []logic.Command{lightCommand}}
	controller := implement.NewRemoteControl(2)
	controller.SetCommand(0, lightCommand)
	controller.Press(1)
	controller.UndoPress()
	controller.SetCommand(1, party)
	controller.Press(1)
	controller.UndoPress()
	controller.UndoPress()
}
