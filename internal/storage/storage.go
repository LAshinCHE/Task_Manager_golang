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
	if _, err := os.Stat(s.filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if errCreateFile := s.createFile(); errCreateFile != nil {
				return errCreateFile
			}
		} else {
			return err
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

	return os.WriteFile(s.filename, newb, 0644)
}

func (s *Storage) TakeTask(taskID uint64, employeeID uint64) error {
	if _, err := os.Stat(s.filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if errCreateFile := s.createFile(); errCreateFile != nil {
				return errCreateFile
			}
		} else {
			return err
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

	for i := 0; i < len(RecordsTask); i++ {
		if RecordsTask[i].ID == uint64(taskID) {
			RecordsTask[i].EmployeeID = employeeID
			RecordsTask[i].Status = models.OnWorck
		}
	}
	newb, err := json.Marshal(RecordsTask)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newb, 0644)
}

func (s *Storage) CompleteTask(taskID uint64) error {
	if _, err := os.Stat(s.filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if errCreateFile := s.createFile(); errCreateFile != nil {
				return errCreateFile
			}
		} else {
			return err
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

	for i := 0; i < len(RecordsTask); i++ {
		if RecordsTask[i].ID == uint64(taskID) {
			RecordsTask[i].Status = models.Completed
		}
	}
	newb, err := json.Marshal(RecordsTask)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newb, 0644)
}

func (s *Storage) ChangeEmployee(taskID uint64, employeeID uint64) error {
	if _, err := os.Stat(s.filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if errCreateFile := s.createFile(); errCreateFile != nil {
				return errCreateFile
			}
		} else {
			return err
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

	for i := 0; i < len(RecordsTask); i++ {
		if RecordsTask[i].ID == uint64(taskID) {
			RecordsTask[i].EmployeeID = employeeID
		}
	}
	newb, err := json.Marshal(RecordsTask)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newb, 0644)
}

func (s *Storage) TaskList() ([]models.TaskDTO, error) {
	if _, err := os.Stat(s.filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if errCreateFile := s.createFile(); errCreateFile != nil {
				return nil, errCreateFile
			}
		} else {
			return nil, err
		}
	}

	b, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var RecordsTask []models.Task

	if err := json.Unmarshal(b, &RecordsTask); err != nil {
		return nil, err
	}

	var TaskDTOs []models.TaskDTO

	for i := 0; i < len(RecordsTask); i++ {
		TaskDTOs = append(TaskDTOs, models.TaskDTO{
			Name:        RecordsTask[i].Name,
			Description: RecordsTask[i].Description,
		})
	}
	return TaskDTOs, nil
}

func (s *Storage) DeleteTask(taskID uint64) error {
	if _, err := os.Stat(s.filename); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if errCreateFile := s.createFile(); errCreateFile != nil {
				return errCreateFile
			}
		} else {
			return err
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

	for i := 0; i < len(RecordsTask); i++ {
		if RecordsTask[i].ID == uint64(taskID) {
			RecordsTask = append(RecordsTask[:i], RecordsTask[i+1:]...)
			i--
		}
	}
	newb, err := json.Marshal(RecordsTask)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, newb, 0644)
}

func (s *Storage) createFile() error {
	f, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
