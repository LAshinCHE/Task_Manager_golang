package cli

const (
	help       = "help"
	addTask    = "add"
	deleteTask = "delete"
	listTask   = "list"
	findTask   = "find"
)

type command struct {
	name        string
	description string
}
