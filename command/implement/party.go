package implement

import (
	"github.com/LFZJun/go-lib/command/logic"
)

type Party struct {
	Commands []logic.Command
}

func (p *Party) Execute() {
	for _, v := range p.Commands {
		v.Execute()
	}
}

func (p *Party) Undo() {
	for _, v := range p.Commands {
		v.Undo()
	}
}
