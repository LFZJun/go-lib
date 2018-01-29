package command

import "errors"

type RemoteControl struct {
	commands    []Command
	lastCommand Command
}

func (rc *RemoteControl) Press(i int) {
	command := rc.commands[i]
	command.Execute()
	rc.lastCommand = command
}

func (rc *RemoteControl) UndoPress() {
	rc.lastCommand.Undo()
}

func (rc *RemoteControl) SetCommand(site int, command Command) error {
	if site >= len(rc.commands) {
		return errors.New("exceed length")
	}
	rc.commands[site] = command
	return nil
}

func NewRemoteControl(commandAmount int) *RemoteControl {
	noCommand := new(NoCommand)
	commands := make([]Command, commandAmount)
	for i := 0; i < commandAmount; i++ {
		commands[i] = noCommand
	}
	return &RemoteControl{commands: commands, lastCommand: noCommand}
}
