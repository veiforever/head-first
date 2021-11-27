package state

import (
	"fmt"
	"testing"
)

func TestGumballMachine(t *testing.T) {
	g := NewGumballMachine(5)
	fmt.Println(g.String())

	g.InsertQuarter()
	g.TurnCrank()

	fmt.Println(g.String())

	g.InsertQuarter()
	g.TurnCrank()
	g.InsertQuarter()
	g.TurnCrank()

	fmt.Println(g.String())
}
