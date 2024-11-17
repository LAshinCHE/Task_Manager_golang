package models

const (
	AtStart   = "at start"
	OnWorck   = "in progress"
	Completed = "completed"
)

type status string

type TaskDTO struct {
	Name        string
	Description string
}

type Task struct {
	ID          uint64
	EmployeeID  uint64
	Name        string
	Description string

	Status status
}
