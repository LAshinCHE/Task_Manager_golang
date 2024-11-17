package errors

import "errors"

var (
	NonCommandError = errors.New(`
		The command line is empty.\n
		Enter the command to interact with the application.
		Enter help to see command list.
	`)

	NotEnoughtArgumentToAddFunctionError = errors.New(`
		Not enought arguments in add function.
		Please enter two arguments name task and task description.
	`)
	NotEnoughtArgumentToTakeTaskFunctionError = errors.New(`
		Not enought arguments in take function.
		Please enter two arguments name task and ID worker ID.
	`)
	NotEnoughtArgumentToCompleteFunctionError = errors.New(`
		Not enought arguments in take function.
		Please enter two arguments name task and ID worker ID.
	`)

	EmptyStorageError = errors.New(`
		Empty storage.
	`)
)
