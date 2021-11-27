package light

type Light struct{}

func (l *Light) On()  { println("light on") }
func (l *Light) Off() { println("light off") }
