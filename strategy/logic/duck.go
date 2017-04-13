package logic

type Duck struct {
	QuackInterface QuackBehavior
}

func (d *Duck) PerformQuack() {
	d.QuackInterface.Quack()
}
