package logic

type Command interface {
	Execute()
	Undo()
}