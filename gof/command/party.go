package command

type Party struct {
	Commands []Command
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
