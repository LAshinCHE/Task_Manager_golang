package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/LAshinCHE/Task_Manager_golang/internal/cli/errors"
	"github.com/LAshinCHE/Task_Manager_golang/internal/models"
)

type service interface {
	Add(models.TaskDTO) error
	TakeTask(taskID uint64, employeeID uint64) error
	CompleteTask(taskID uint64) error
	ChangeEmployee(taskID uint64, employeeID uint64) error
	DeleteTask(taskID uint64) error
	TaskList() ([]models.TaskDTO, error)
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
	case takeTask:
		return c.takeTask(args[1:])
	case completeTask:
		return c.completeTask(args[1:])
	case changeEmployee:
		return c.changeEmployee(args[1:])
	case listTask:
		return c.listTask()
	case deleteTask:
		return c.deleteTask(args[1:])
	default:
		c.Help()
		return errors.NonCommandError
	}
}

func (c *CLI) Help() {
	commandList := emplymentCommandList()

	for _, com := range commandList {
		fmt.Printf("Command Name: %s\nCommand Description: %s\n\n", com.name, com.description)
	}
}

func (c *CLI) Add(args []string) error {
	if len(args) < 2 {
		return errors.NotEnoughtArgumentToAddFunctionError
	}

	name := args[0]
	description := strings.Join(args[1:], " ")

	task := models.TaskDTO{
		Name:        name,
		Description: description,
	}

	return c.service.Add(task)
}
func (c *CLI) takeTask(args []string) error {
	if len(args) != 2 {
		return errors.NotEnoughtArgumentToAddFunctionError
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	workerID, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	return c.service.TakeTask(uint64(taskID), uint64(workerID))
}

func (c *CLI) completeTask(args []string) error {
	if len(args) != 1 {
		return errors.NotEnoughtArgumentToCompleteFunctionError
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	return c.service.CompleteTask(uint64(taskID))
}

func (c *CLI) changeEmployee(args []string) error {
	if len(args) != 2 {
		return errors.NotEnoughtArgumentToAddFunctionError
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	workerID, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	return c.service.ChangeEmployee(uint64(taskID), uint64(workerID))
}

func (c *CLI) listTask() error {

	tasks, err := c.service.TaskList()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return errors.EmptyStorageError
	}

	for i := 0; i < len(tasks); i++ {
		fmt.Printf("Task# %s\tDescription: %s", tasks[i].Name, tasks[i].Description)
	}
	return nil
}

func (c *CLI) deleteTask(args []string) error {
	if len(args) != 2 {
		return errors.NotEnoughtArgumentToAddFunctionError
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	return c.service.DeleteTask(uint64(taskID))
}
