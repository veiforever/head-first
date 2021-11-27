package state

var _ State = &NoQuarterState{}

type NoQuarterState struct {
	gumballMachine *gumballMachine
}

func (s *NoQuarterState) String() string {
	return "waitting for quarter"
}

func (ns *NoQuarterState) InsertQuarter() {
	println("You inserted a quarter")
	ns.gumballMachine.setState(ns.gumballMachine.getState(STATEHASQUARTER))
}

func (ns *NoQuarterState) EjectQuarter() {
	println("You haven't inserted a quarter")
}

func (ns *NoQuarterState) TurnCrank() {
	println("You turned, but therer 's no quarter")
}

func (ns *NoQuarterState) Dispense() {
	println("You need to pay first")
}
