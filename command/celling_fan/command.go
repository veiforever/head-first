package cellingfan

type CellingFanHighCommand struct {
	fan       *CellingFan
	prevSpeed int
}

func NewCellingFanHighCommand(fan *CellingFan) *CellingFanHighCommand {
	return &CellingFanHighCommand{fan: fan}
}

func (c *CellingFanHighCommand) Execute() {
	c.prevSpeed = c.fan.GetSpeed()
	c.fan.High()
}
func (c *CellingFanHighCommand) Undo() {
	switch c.prevSpeed {
	case MEDIUM:
		c.fan.Medium()
	case LOW:
		c.fan.Low()
	case OFF:
		c.fan.Off()
	}
}

type CellingFanOffCommand struct {
	fan       *CellingFan
	prevSpeed int
}

func NewCellingFanOffCommand(fan *CellingFan) *CellingFanOffCommand {
	return &CellingFanOffCommand{fan: fan}
}

func (c *CellingFanOffCommand) Execute() {
	c.prevSpeed = c.fan.GetSpeed()
	c.fan.Off()
}

func (c *CellingFanOffCommand) Undo() {
	switch c.prevSpeed {
	case MEDIUM:
		c.fan.Medium()
	case LOW:
		c.fan.Low()
	case HIGH:
		c.fan.High()
	}
}

type CellingFanMediumCommand struct {
	fan       *CellingFan
	prevSpeed int
}

func NewCellingFanMediumCommand(fan *CellingFan) *CellingFanMediumCommand {
	return &CellingFanMediumCommand{fan: fan}
}

func (c *CellingFanMediumCommand) Execute() {
	c.prevSpeed = c.fan.GetSpeed()
	c.fan.Medium()
}

func (c *CellingFanMediumCommand) Undo() {
	switch c.prevSpeed {
	case OFF:
		c.fan.Off()
	case LOW:
		c.fan.Low()
	case HIGH:
		c.fan.High()
	}
}
