package service

import "github.com/LAshinCHE/Task_Manager_golang/internal/models"

type storage interface {
	Add(models.TaskDTO) error
	TakeTask(uint64, uint64) error
	CompleteTask(uint64) error
	ChangeEmployee(uint64, uint64) error
	DeleteTask(uint64) error
	TaskList() ([]models.TaskDTO, error)
}

type Service struct {
	storage storage
}

func NewService(st storage) *Service {
	return &Service{
		storage: st,
	}
}

func (s Service) Add(t models.TaskDTO) error {
	return s.storage.Add(t)
}

func (s Service) TakeTask(taskID, employeeID uint64) error {
	return s.storage.TakeTask(taskID, employeeID)
}

func (s Service) CompleteTask(taskID uint64) error {
	return s.storage.CompleteTask(taskID)
}

func (s Service) ChangeEmployee(taskID, employeeID uint64) error {
	return s.storage.ChangeEmployee(taskID, employeeID)
}

func (s Service) DeleteTask(taskID uint64) error {
	return s.storage.DeleteTask(taskID)
}

func (s Service) TaskList() ([]models.TaskDTO, error) {
	return s.storage.TaskList()
}
