package control

import (
	"testing"

	cellingfan "github.com/projectred/head—first/command/celling_fan"
	garagedoor "github.com/projectred/head—first/command/garage_door"
	"github.com/projectred/head—first/command/light"
)

func TestControl(t *testing.T) {
	remoteControl := NewRemoteControl()

	l := &light.Light{}
	gd := &garagedoor.GarageDoor{}
	fan := cellingfan.NewCellingFan("living room")

	remoteControl.SetCommand(0, light.NewLightOnCommand(l), light.NewLightOffCommand(l))
	remoteControl.SetCommand(1, garagedoor.NewGarageDoorOpenCommand(gd), garagedoor.NewGarageDoorCloseCommand(gd))
	remoteControl.SetCommand(2, cellingfan.NewCellingFanHighCommand(fan), cellingfan.NewCellingFanOffCommand(fan))
	remoteControl.SetCommand(3, cellingfan.NewCellingFanMediumCommand(fan), cellingfan.NewCellingFanOffCommand(fan))

	println(remoteControl.String())

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	println(remoteControl.String())
	remoteControl.UndoButtonWasPushed()

	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(0)
	println(remoteControl.String())
	remoteControl.UndoButtonWasPushed()

	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
	println(remoteControl.String())
	remoteControl.UndoButtonWasPushed()

	remoteControl.OnButtonWasPushed(1)
	println(remoteControl.String())
	remoteControl.UndoButtonWasPushed()
}
