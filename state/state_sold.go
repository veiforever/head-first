package state

var _ State = &SoldState{}

type SoldState struct {
	gumballMachine *gumballMachine
}

func (s *SoldState) String() string {
	return "sold"
}

func (ns *SoldState) InsertQuarter() {
	println("Please wait, we're already giving you a gumall")
}

func (ns *SoldState) EjectQuarter() {
	println("Sorry, you already turned the crank")
}

func (ns *SoldState) TurnCrank() {
	println("Turning twice doesn't get you another gumball!")
}

func (ns *SoldState) Dispense() {
	ns.gumballMachine.ReleaseBall()
	if ns.gumballMachine.GetCount() > 0 {
		ns.gumballMachine.setState(ns.gumballMachine.getState(STATENOQUARTER))
	} else {
		println("Oops, out of gumballs!")
		ns.gumballMachine.setState(ns.gumballMachine.getState(STATESOLDOUT))
	}
}
