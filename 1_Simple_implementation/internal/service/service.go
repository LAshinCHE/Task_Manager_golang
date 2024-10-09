package service

import "github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"

type storage interface {
	Add(model.TaksInput) error
	Get(string, ...int) ([]model.Task, error)
	Update(model.Task) error
	Delete(int) error
}
