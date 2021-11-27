package command

type Command interface {
	Execute()
	Undo()
}

type NoCommand struct{}

func (c *NoCommand) Execute() {}
func (c *NoCommand) Undo()    {}
