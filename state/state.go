package state

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
	String() string
}
