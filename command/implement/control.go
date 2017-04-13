package implement

import (
	"github.com/LFZJun/go-lib/command/logic"
	"qiniupkg.com/x/errors.v7"
)

type RemoteControl struct {
	commands    []logic.Command
	lastCommand logic.Command
}

func (rc *RemoteControl) Press(i int) {
	command := rc.commands[i]
	command.Execute()
	rc.lastCommand = command
}

func (rc *RemoteControl) UndoPress() {
	rc.lastCommand.Undo()
}

func (rc *RemoteControl) SetCommand(site int, command logic.Command) error {
	if site >= len(rc.commands) {
		return errors.New("exceed length")
	}
	rc.commands[site] = command
	return nil
}

func NewRemoteControl(commandAmount int) *RemoteControl {
	noCommand := new(NoCommand)
	commands := make([]logic.Command, commandAmount)
	for i := 0; i < commandAmount; i++ {
		commands[i] = noCommand
	}
	return &RemoteControl{commands: commands, lastCommand: noCommand}
}
