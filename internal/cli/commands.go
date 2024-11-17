package cli

const (
	help           = "help"
	addTask        = "add"
	takeTask       = "take"
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
		{
			name:        takeTask,
			description: "Take task to work",
		},
		{
			name:        completeTask,
			description: "Take task to work",
		},
		{
			name:        changeEmployee,
			description: "Change employee in selected task",
		},
		{
			name:        listTask,
			description: "Print task list in CLI",
		},
		{
			name:        deleteTask,
			description: "Delete task from storage",
		},
	}
	return commands
}
