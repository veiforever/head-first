package garagedoor

type GarageDoorOpen struct {
	garageDoor *GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor *GarageDoor) *GarageDoorOpen {
	return &GarageDoorOpen{garageDoor: garageDoor}
}

func (c *GarageDoorOpen) Execute() { c.garageDoor.On() }
func (c *GarageDoorOpen) Undo()    { c.garageDoor.Off() }

type GarageDoorClose struct {
	garageDoor *GarageDoor
}

func NewGarageDoorCloseCommand(garageDoor *GarageDoor) *GarageDoorClose {
	return &GarageDoorClose{garageDoor: garageDoor}
}

func (c *GarageDoorClose) Execute() { c.garageDoor.Off() }

func (c *GarageDoorClose) Undo() { c.garageDoor.On() }
