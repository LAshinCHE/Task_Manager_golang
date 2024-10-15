package service

import (
	"fmt"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"
)

type storage interface {
	AddTask(model.TaksInput) error
	Get(string, ...int) ([]model.Task, error)
	Update(model.Task) error
	Delete(int) error
}

type Service struct {
	storage storage
}

func NewService(st storage) *Service {
	return &Service{
		storage: st,
	}
}

func (s Service) AddTask(excecutionPeriod int, descption string)

func (s *Service) Help() {
	fmt.Println("-help - print list command")
}
