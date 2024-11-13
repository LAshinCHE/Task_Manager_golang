package cli

type service interface{}

type CLI struct {
	service service
}

func NewCLI(serv service) *CLI {
	return &CLI{
		service: serv,
	}
}

func (c *CLI) Execute(args []string) error {
	return nil
}
