package strategy

type Duck struct {
	QuackInterface QuackBehavior
}

func (d *Duck) PerformQuack() {
	d.QuackInterface.Quack()
}

func NewDuck(quack QuackBehavior) *Duck {
	return &Duck{QuackInterface: quack}
}