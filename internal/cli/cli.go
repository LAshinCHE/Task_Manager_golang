package cli

import (
	"fmt"
	"strings"

	"github.com/LAshinCHE/Task_Manager_golang/internal/cli/errors"
	"github.com/LAshinCHE/Task_Manager_golang/internal/models"
)

type service interface {
	Add(models.TaskDTO) error
}

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
	case help:
		c.Help()
		return nil
	case addTask:
		return c.Add(args[1:])
	}
	return nil
}

func (c *CLI) Help() {
	commandList := emplymentCommandList()

	for _, com := range commandList {
		fmt.Printf("Command Name: %s\nCommand Description: %s\n\n", com.name, com.description)
	}
}

func (c *CLI) Add(args []string) error {
	if len(args) < 2 {
		return errors.NotEnoughtArgumentToAddFunction
	}

	name := args[0]
	description := strings.Join(args[1:], " ")

	task := models.TaskDTO{
		Name:        name,
		Description: description,
	}

	return c.service.Add(task)
}
