package storage

import "github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"

type Storage struct {
	filename string
}

func NewStorage(f string) Storage {
	return Storage{
		filename: f,
	}
}

func (s Storage) Add(model.TaksInput) error {
	return nil
}

func (s Storage) Get(string, ...int) ([]model.Task, error) {
	return nil, nil
}

func (s Storage) Update(model.Task) error {
	return nil
}

func (s Storage) Delete(int) error {
	return nil
}
