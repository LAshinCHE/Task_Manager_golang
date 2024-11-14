package errors

import "errors"

var (
	NonCommandError = errors.New(`
		The command line is empty.\n
		Enter the command to interact with the application.\n
		Enter help to see command list.\n
		`)
)
