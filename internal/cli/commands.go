package cli

const (
	help           = "help"
	addTask        = "add"
	tekeTask       = "take"
	completeTask   = "complete"
	changeEmployee = "change"
	listTask       = "list"
	deleteTask     = "delete"
)

type command struct {
	name        string
	description string
}

func emplymentCommandList() []command {
	commands := []command{
		{
			name:        help,
			description: "Help function. Prints a list of commands and their description.",
		},
		{
			name:        addTask,
			description: "Add taks to storage.",
		},
	}
	return commands
}
