package storage

import (
	"encoding/json"
	"errors"
	"fmt"
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

	var taskRecords []TaskRecord
	if errUnmarshal := json.Unmarshal(b, &taskRecords); errUnmarshal != nil {
		return errUnmarshal
	}
	taskRec := TaskRecord{
		ID:          len(taskRecords),
		Name:        string(task.Name),
		Description: string(task.Description),
		CreatedAt:   time.Now(),
	}
	taskRecords = append(taskRecords, taskRec)
	newB, err := json.MarshalIndent(taskRecords, " ", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newB, 0666)
}

func (s Storage) createFile() error {
	f, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (s Storage) ListTask() ([]model.TaskDTO, error) {
	if _, err := os.Stat(s.filename); errors.Is(err, os.ErrNotExist) {
		if errCreateFile := s.createFile(); errCreateFile != nil {
			return nil, errCreateFile
		}
	}

	b, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var taskRecords []TaskRecord

	if err := json.Unmarshal(b, &taskRecords); err != nil {
		return nil, err
	}

	tasksTransfer := make([]model.TaskDTO, len(taskRecords))

	for i := 0; i < len(taskRecords); i++ {
		tasksTransfer[i] = model.TaskDTO{
			Name:        model.Name(taskRecords[i].Name),
			Description: model.Description(taskRecords[i].Description),
		}
	}
	return tasksTransfer, nil
}

func (s Storage) DeleteTask(deletedTask model.TaskDTO) error {
	if _, err := os.Stat(s.filename); errors.Is(err, os.ErrNotExist) {
		if errCreateFile := s.createFile(); errCreateFile != nil {
			return errCreateFile
		}
	}

	b, err := os.ReadFile(s.filename)

	if err != nil {
		return err
	}

	var taskRecords []TaskRecord

	if err := json.Unmarshal(b, &taskRecords); err != nil {
		return err
	}

	if len(taskRecords) == 0 {
		return fmt.Errorf("Task list is empty, please add some task in storage")
	}

	found := false
	for i := 0; i < len(taskRecords); i++ {
		if taskRecords[i].Name == string(deletedTask.Name) {
			taskRecords = append(taskRecords[:i], taskRecords[i+1:]...)
			found = true
		}
	}
	if found == false {
		return fmt.Errorf("Task was not found in storage")
	}
	var byteNewTaskREcods []byte
	byteNewTaskREcods, err = json.MarshalIndent(taskRecords, " ", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, byteNewTaskREcods, 0666)
}
