package control

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/projectred/head—first/command"
)

type RemoteControl struct {
	onCommands  []command.Command
	offCommands []command.Command
	undoCommand command.Command
}

func NewRemoteControl() *RemoteControl {
	noCommand := &command.NoCommand{}
	c := &RemoteControl{make([]command.Command, 7), make([]command.Command, 7), noCommand}
	for i := 0; i < 7; i++ {
		c.onCommands[i] = noCommand
		c.offCommands[i] = noCommand
	}
	return c
}

func (c *RemoteControl) SetCommand(slot int, onCommand, offCommand command.Command) {
	c.onCommands[slot] = onCommand
	c.offCommands[slot] = offCommand
}

func (c *RemoteControl) OnButtonWasPushed(slot int) {
	c.onCommands[slot].Execute()
	c.undoCommand = c.onCommands[slot]
}

func (c *RemoteControl) OffButtonWasPushed(slot int) {
	c.offCommands[slot].Execute()
	c.undoCommand = c.offCommands[slot]
}

func (c *RemoteControl) UndoButtonWasPushed() { c.undoCommand.Undo() }

func (c *RemoteControl) String() string {
	builder := strings.Builder{}
	builder.WriteString("\n------ Remote Control --------\n")
	for i := 0; i < len(c.onCommands); i++ {
		builder.WriteString("[slot " + strconv.Itoa(i) + "]")
		for _, command := range []command.Command{c.onCommands[i], c.offCommands[i]} {
			builder.WriteString("\t")
			t := reflect.TypeOf(command)
			for t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			builder.WriteString(t.Name())
		}
		builder.WriteString("\n")
	}
	for undoType := reflect.TypeOf(c.undoCommand); ; {
		if undoType.Kind() != reflect.Ptr {
			builder.WriteString("[undo]\t" + undoType.Name())
			break
		}
		undoType = undoType.Elem()
	}
	return builder.String()
}
