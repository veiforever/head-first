package state

import (
	"math/rand"
	"time"
)

var _ State = &HasQuarterState{}

type HasQuarterState struct {
	gumballMachine *gumballMachine
}

func (s *HasQuarterState) String() string {
	return "having quarter"
}

func (ns *HasQuarterState) InsertQuarter() {
	println("You can't insett another quarter")
}

func (ns *HasQuarterState) EjectQuarter() {
	println("Quarter returned")
	ns.gumballMachine.setState(ns.gumballMachine.getState(STATENOQUARTER))
}

func (ns *HasQuarterState) TurnCrank() {
	println("You turned...")
	if rand.Intn(10) == 0 && ns.gumballMachine.GetCount() > 1 {
		ns.gumballMachine.setState(ns.gumballMachine.getState(STATEWINNER))
	} else {
		ns.gumballMachine.setState(ns.gumballMachine.getState(STATESOLD))
	}
}

func (ns *HasQuarterState) Dispense() {
	println("No gumball dispensed")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
