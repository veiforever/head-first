package state

import "fmt"

const (
	STATESOLDOUT = iota
	STATENOQUARTER
	STATEHASQUARTER
	STATESOLD
	STATEWINNER
	STATEMAX
)

type gumballMachine struct {
	states []State
	state  State
	count  int
}

func NewGumballMachine(numberGumballs int) *gumballMachine {
	gumballMachine := &gumballMachine{}
	gumballMachine.states = []State{&SoldOutState{gumballMachine}, &NoQuarterState{gumballMachine}, &HasQuarterState{gumballMachine}, &SoldState{gumballMachine}, &WinnerState{gumballMachine}}
	gumballMachine.count = numberGumballs
	if numberGumballs > 0 {
		gumballMachine.state = gumballMachine.states[STATENOQUARTER]
	}
	return gumballMachine
}

func (gm *gumballMachine) setState(state State) { gm.state = state }

func (gm *gumballMachine) getState(state int) State { return gm.states[state] }

func (gm *gumballMachine) InsertQuarter() { gm.state.InsertQuarter() }

func (gm *gumballMachine) EjectQuarter() { gm.state.EjectQuarter() }

func (gm *gumballMachine) TurnCrank() { gm.state.TurnCrank(); gm.state.Dispense() }

func (gm *gumballMachine) ReleaseBall() {
	println("A gumball comes rolling out the slot...")
	if gm.count > 0 {
		gm.count--
	}
}

func (gm *gumballMachine) GetCount() int { return gm.count }

func (gm *gumballMachine) String() string {
	return fmt.Sprintf("Inventory: %d gumballs\nMachine is %s", gm.count, gm.state.String())
}
