package service

import "github.com/LAshinCHE/Task_Manager_golang/internal/models"

type storage interface {
	Add(models.TaskDTO) error
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
