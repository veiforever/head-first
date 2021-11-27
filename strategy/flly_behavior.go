package strategy

type FilyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (f *FlyWithWings) Fly() { println("I'm flying!") }

type FlyNoWay struct{}

func (f *FlyNoWay) Fly() { println("I can't fly") }

type FlyRocketPowered struct{}

func (f *FlyRocketPowered) Fly() { println("I'm flying with a rocket") }
