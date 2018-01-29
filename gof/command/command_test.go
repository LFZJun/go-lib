package command

import "testing"

func TestCommand(t *testing.T) {
	light := &Light{Site: "hall"}
	lightCommand := &LightCommand{Light: light}
	party := &Party{Commands: []Command{lightCommand}}
	controller := NewRemoteControl(2)
	controller.SetCommand(0, lightCommand)
	controller.Press(1)
	controller.UndoPress()
	controller.SetCommand(1, party)
	controller.Press(1)
	controller.UndoPress()
	controller.UndoPress()
}
