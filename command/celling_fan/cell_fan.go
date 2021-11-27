package cellingfan

import "strconv"

const (
	OFF = iota
	LOW
	MEDIUM
	HIGH
)

type CellingFan struct {
	location string
	speed    int
}

func NewCellingFan(locaiton string) *CellingFan { return &CellingFan{location: locaiton, speed: OFF} }

func (c *CellingFan) High()         { c.speed = HIGH; println(c.String()) }
func (c *CellingFan) Medium()       { c.speed = MEDIUM; println(c.String()) }
func (c *CellingFan) Low()          { c.speed = LOW; println(c.String()) }
func (c *CellingFan) Off()          { c.speed = OFF; println(c.String()) }
func (c *CellingFan) GetSpeed() int { return c.speed }

func (c *CellingFan) String() string {
	return c.location + " celling fan is on speed" + strconv.Itoa(c.GetSpeed())
}
