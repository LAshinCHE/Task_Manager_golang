package cli

import (
	"flag"
	"fmt"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"
)

type Service interface {
	AddTask(model.TaskDTO) error
	ListTask() ([]model.TaskDTO, error)
	DeleteTask(model.TaskDTO) error
}

type Deps struct {
	Service Service
}

type CLI struct {
	Deps
	commandList []command
}

func NewCLI(d Deps) CLI {
	return CLI{
		Deps: d,
		commandList: []command{
			{
				name:        help,
				description: "Справка",
			},
			{
				name:        addTask,
				description: "Добавляет задачу в хранилище: использование add --name=SomeName --description=some_task_description",
			},
			{
				name:        deleteTask,
				description: "Удаляет задачу с  хранилища: использование delete --name=SomeName ",
			},
			{
				name:        listTask,
				description: "Показывает список задач: использование list ",
			},
			{
				name:        findTask,
				description: "Ищет задачу по ее имени: find --name=SomeName ",
			},
		},
	}
}

func (c CLI) HandleCommand(commands []string) error {
	if len(commands) == 0 {
		c.help()
		return fmt.Errorf("command isn't set, u can choose one from command list")
	}
	commandName := commands[0]
	switch commandName {
	case help:
		c.help()
		return nil
	case addTask:
		return c.addTask(commands[1:])
	case deleteTask:
		return c.deleteTask(commands[1:])
	case listTask:
		return c.listTask()
	case findTask:
		return c.findTask()
	}

	return nil
}

func (c CLI) help() {
	fmt.Println("Command list:")
	for _, command := range c.commandList {
		fmt.Println(command.name, command.description)
	}
}

// TODO addTask - функция должна добавлять задачу в хранилище
// TODO - описать DTO передоваемые между модулями данные
func (c CLI) addTask(args []string) error {
	var name, description string

	fs := flag.NewFlagSet(addTask, flag.ContinueOnError)
	fs.StringVar(&name, "name", "", "use --name=someName")
	fs.StringVar(&description, "description", "", "use --description=some_description")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if len(name) == 0 {
		return fmt.Errorf("Your task has empty name, please enter the name with flag --name")
	}

	if len(description) == 0 {
		return fmt.Errorf("Your task has empty description, please enter the name with flag --description")
	}

	task := model.TaskDTO{
		Name:        model.Name(name),
		Description: model.Description(description),
	}

	return c.Service.AddTask(task)
}

// TODO deleteTask - функция должна удалять задачу из хранилище
func (c CLI) deleteTask(args []string) error {
	var name string

	fs := flag.NewFlagSet(deleteTask, flag.ContinueOnError)

	fs.StringVar(&name, "name", "", "use --name=someName")

	if err := fs.Parse(args); err != nil {
		return nil
	}

	if len(name) == 0 {
		return fmt.Errorf("Your task has empty name, please enter the name with flag --name")
	}

	task := model.TaskDTO{
		Name: model.Name(name),
	}

	return c.Service.DeleteTask(task)
}

// TODO findTask - функция должна находить задачу в хранилище
func (c CLI) findTask() error {
	return nil
}

// TODO listTask - функция должна показывать список текущих задач
func (c CLI) listTask() error {
	tasks, err := c.Service.ListTask()

	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return fmt.Errorf("Tasks storage is empty")
	}

	fmt.Println("Tasks list")
	for i := 0; i < len(tasks); i++ {
		fmt.Printf("Name: %s Description: %s \n", tasks[i].Name, tasks[i].Description)
	}

	return nil
}
