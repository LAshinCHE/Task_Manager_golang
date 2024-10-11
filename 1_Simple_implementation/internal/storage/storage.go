package storage

import (
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"
)

const storageFile = "../data/data.json"

type Storage struct {
	storage *os.File
}

func NewStorage() (Storage, error) {
	file, err := os.OpenFile(storageFile, os.O_CREATE, 0777)
	if err != nil {
		return Storage{}, err
	}
	return Storage{storage: file}, nil
}

func (s Storage) CloseStorage() error {
	return s.storage.Close()
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
