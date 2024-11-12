package model

type Name string
type Description string

// type Task struct {
// 	ID              int
// 	EmployeeID      int
// 	AppointmentTime time.Time
// 	Description     string

// 	IsCompleted     bool
// 	ExecutionPeriod time.Time

// 	IssueTime time.Time
// }

type TaskDTO struct {
	Name        Name
	Description Description
}
