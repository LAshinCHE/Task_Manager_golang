package cli

import (
	"fmt"
)

type Service interface {
	AddTask(excecutionPeriod int, descption string) error
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
		return c.addTask()
	case deleteTask:
		return c.deleteTask()
	case listTask:
		c.listTask()
		return nil
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
func (c CLI) addTask() error {
	return nil
}

// TODO deleteTask - функция должна удалять задачу из хранилище
func (c CLI) deleteTask() error {
	return nil
}

// TODO findTask - функция должна находить задачу в хранилище
func (c CLI) findTask() error {
	return nil
}

// TODO listTask - функция должна показывать список текущих задач
func (c CLI) listTask() {
}
