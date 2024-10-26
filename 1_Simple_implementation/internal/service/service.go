package service

import (
	"fmt"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"
)

type storage interface {
	AddTask(model.TaskDTO) error
	ListTask() ([]model.TaskDTO, error)
	DeleteTask(model.TaskDTO) error
	FindTask(model.TaskDTO) (*model.TaskDTO, error)
	//Get() ([]model.TaskDTO, error)
	// Update(model.Task) error
	// Delete(int) error
}

type Service struct {
	storage storage
}

func NewService(st storage) *Service {
	return &Service{
		storage: st,
	}
}

func (s Service) AddTask(taskAdded model.TaskDTO) error {
	return s.storage.AddTask(taskAdded)
}

func (s Service) DeleteTask(taskDeleted model.TaskDTO) error {
	return s.storage.DeleteTask(taskDeleted)
}

func (s *Service) Help() {
	fmt.Println("-help - print list command")
}

func (s *Service) ListTask() ([]model.TaskDTO, error) {
	return s.storage.ListTask()
}

func (s *Service) FindTask(taskToBeFound model.TaskDTO) (*model.TaskDTO, error) {
	return s.storage.FindTask(taskToBeFound)
}
