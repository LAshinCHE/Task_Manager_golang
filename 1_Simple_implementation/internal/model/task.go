package model

import "time"

type Task struct {
	ID              int
	EmployeeID      int
	AppointmentTime time.Time
	Description     string

	IsCompleted     bool
	ExecutionPeriod time.Time

	IssueTime time.Time
}

type TaksInput struct {
	ID          int
	EmployeeID  int
	Description string
}
