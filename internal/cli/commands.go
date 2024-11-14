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
