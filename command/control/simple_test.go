package control

import (
	"testing"

	garagedoor "github.com/projectred/head—first/command/garage_door"
	"github.com/projectred/head—first/command/light"
)

func TestSimpleRemoteControl(t *testing.T) {
	remote := &SimpleRemoteControl{}
	lightOn := light.NewLightOnCommand(&light.Light{})
	garageDoorOpen := garagedoor.NewGarageDoorOpenCommand(&garagedoor.GarageDoor{})
	remote.SetCommand(lightOn)
	remote.ButtonWasPassed()

	remote.SetCommand(garageDoorOpen)
	remote.ButtonWasPassed()
}
