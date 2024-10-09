package comands

type Commander struct {
	arguments []string
}

func NewCommander(args []string) *Commander {
	return &Commander{
		arguments: args,
	}
}

func (c *Commander) HandleCommand()
