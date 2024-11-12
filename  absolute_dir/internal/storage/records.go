package storage

import (
	"time"
)

type TaskRecord struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`

	// EmployeeID  int       `json:"employee_id"`

	// IsCompleted     bool
	// ExecutionPeriod time.Time

	// IssueTime time.Time
}

// func transform(task model.TaskDTO) TaskRecord {
// 	return TaskRecord{
// 		ID:
// 	}
// }
