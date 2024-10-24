package storage

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/model"
)

const storageFile = "./data/data.json"

type Storage struct {
	filename string
}

func NewStorage(fln string) Storage {
	return Storage{
		filename: fln,
	}
}

func (s Storage) AddTask(task model.TaskDTO) error {
	if _, err := os.Stat(s.filename); errors.Is(err, os.ErrNotExist) {
		if errCreateFile := s.createFile(); errCreateFile != nil {
			return errCreateFile
		}
	}

	b, err := os.ReadFile(s.filename)
	if err != nil {
		return err
	}

	var tasks []TaskRecord
	if errUnmarshal := json.Unmarshal(b, &tasks); errUnmarshal != nil {
		return errUnmarshal
	}
	taskRec := TaskRecord{
		ID:          len(tasks),
		Name:        string(task.Name),
		Description: string(task.Description),
		CreatedAt:   time.Now(),
	}
	tasks = append(tasks, taskRec)
	newB, err := json.MarshalIndent(tasks, " ", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newB, 0666)
}

// func (s Storage) Get(string, ...int) ([]model.Task, error) {

// 	return nil, nil
// }

// func (s Storage) Update(model.TaskDTO) error {
// 	return nil
// }

//	func (s Storage) Delete(int) error {
//		return nil
//	}
func (s Storage) createFile() error {
	f, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
