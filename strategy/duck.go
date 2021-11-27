package strategy

type Duck struct {
	quackBehavior QuackBehavior
	filyBehavior  FilyBehavior
}

func (d *Duck) display() {}

func (d *Duck) PerformFly() {
	d.filyBehavior.Fly()
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.Quack()
}
func (d *Duck) Swim() {
	println("All ducks float, event decoys!")
}

func (d *Duck) SetFlyBehavior(fb FilyBehavior)    { d.filyBehavior = fb }
func (d *Duck) SetQuackBehavior(qb QuackBehavior) { d.quackBehavior = d.quackBehavior }
