package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/internal/models"
)

type Storage struct {
	filename string
}

func NewStorage(fn string) *Storage {
	return &Storage{
		filename: fn,
	}
}

func (s *Storage) Add(t models.TaskDTO) error {
	if _, err := os.Stat(s.filename); errors.Is(err, os.ErrNotExist) {
		if errCreateFile := s.createFile(); errCreateFile != nil {
			return errCreateFile
		}
	}

	b, err := os.ReadFile(s.filename)
	if err != nil {
		return err
	}

	var RecordsTask []models.Task

	if err := json.Unmarshal(b, &RecordsTask); err != nil {
		return err
	}

	newTask := models.Task{
		ID:          uint64(len(RecordsTask)),
		Name:        t.Name,
		Description: t.Description,
		Status:      models.AtStart,
	}

	RecordsTask = append(RecordsTask, newTask)
	newb, err := json.Marshal(RecordsTask)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newb, 06666)
}

func (s *Storage) createFile() error {
	f, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
