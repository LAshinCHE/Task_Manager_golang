package cli

import (
	"fmt"

	"github.com/LAshinCHE/Task_Manager_golang/internal/cli/errors"
)

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
	if len(args) == 0 {
		return errors.NonCommandError
	}
	command := args[0]
	switch command {
	case "help":
		c.Help()
		return nil
	}
	return nil
}

func (c *CLI) Help() {
	commandList := []command{
		{
			name:        help,
			description: "Help function. Prints a list of commands and their description.",
		},
	}

	for _, com := range commandList {
		fmt.Printf("Command Name: %s\nCommand Description: %s\n\n", com.name, com.description)
	}
}
