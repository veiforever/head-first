package control

import "github.com/projectred/head—first/command"

type SimpleRemoteControl struct {
	slot command.Command
}

func (c *SimpleRemoteControl) SetCommand(command command.Command) {
	c.slot = command
}

func (c *SimpleRemoteControl) ButtonWasPassed() {
	c.slot.Execute()
}
