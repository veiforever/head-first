package strategy

import "testing"

type MallardDuck struct {
	*Duck
}

func (d *MallardDuck) display() { println("I'm a real mallard duck") }

func NewMallarDuck() *MallardDuck { return &MallardDuck{&Duck{&Quack{}, &FlyWithWings{}}} }

type ModelDuck struct {
	*Duck
}

func (d *ModelDuck) display() { println("I'm a model duck") }

func NewModelDuck() *ModelDuck { return &ModelDuck{&Duck{&Quack{}, &FlyNoWay{}}} }

func TestDuck(t *testing.T) {
	mallard := NewMallarDuck()
	mallard.display()
	mallard.PerformQuack()
	mallard.PerformFly()

	model := NewModelDuck()
	model.display()
	model.PerformFly()
	model.SetFlyBehavior(&FlyRocketPowered{})
	model.PerformFly()

}
