package garagedoor

type GarageDoor struct{}

func (l *GarageDoor) On()  { println("Garage Door is Open") }
func (l *GarageDoor) Off() { println("Garage Door is Close") }
