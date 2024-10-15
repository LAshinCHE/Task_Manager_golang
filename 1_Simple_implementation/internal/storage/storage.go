package storage

import (
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"
)

const storageFile = "./data/data.json"

type Storage struct {
	filename string
}

func NewStorage(name string) Storage {
	return Storage{filename: name}
}

func (s Storage) createFile() error {
	f, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (s Storage) AddTask(model.TaksInput) error {
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
