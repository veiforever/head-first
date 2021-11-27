package state

var _ State = &WinnerState{}

type WinnerState struct {
	gumballMachine *gumballMachine
}

func (s *WinnerState) String() string {
	return "winner"
}

func (ns *WinnerState) InsertQuarter() {
	println("Please wait, we're already iiving you a gumall")
}

func (ns *WinnerState) EjectQuarter() {
	println("Sorry, you already turned the crank")
}

func (ns *WinnerState) TurnCrank() {
	println("Turning twice doesn't get you another gumball!")
}

func (ns *WinnerState) Dispense() {
	println("YOU'RE A WINNER! You get two gumballs for your quarter")
	ns.gumballMachine.ReleaseBall()
	if ns.gumballMachine.GetCount() == 0 {
		ns.gumballMachine.setState(ns.gumballMachine.getState(STATESOLDOUT))
	} else {
		ns.gumballMachine.ReleaseBall()
		if ns.gumballMachine.GetCount() > 0 {
			ns.gumballMachine.setState(ns.gumballMachine.getState(STATENOQUARTER))
		} else {
			println("Oops, out of gumballs!")
			ns.gumballMachine.setState(ns.gumballMachine.getState(STATESOLDOUT))
		}
	}
}
