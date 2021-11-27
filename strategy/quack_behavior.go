package strategy

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() { println("Quack") }

type MuteQuack struct{}

func (q *MuteQuack) Quack() { println("<< Silence >>") }

type Squeak struct{}

func (q *Squeak) Quack() { println("Squeak") }
