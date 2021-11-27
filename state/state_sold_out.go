package state

var _ State = &SoldOutState{}

type SoldOutState struct {
	gumballMachine *gumballMachine
}

func (s *SoldOutState) String() string {
	return "sold out"
}

func (ns *SoldOutState) InsertQuarter() {
	println("You can't insert a quarter, the machine is sould out")
}

func (ns *SoldOutState) EjectQuarter() {
	println("You can't eject, you haven;t inserted a quarter yet")
}

func (ns *SoldOutState) TurnCrank() {
	println("You turned, but there are no gumballs")
}

func (ns *SoldOutState) Dispense() {
	println("No gumabll dispensed")
}
