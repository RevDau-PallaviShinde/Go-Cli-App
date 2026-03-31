package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/models"
)

const taskFile = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	if _, err := os.Stat(taskFile); errors.Is(err, os.ErrNotExist) {
		return []models.Task{}, nil
	}

	file, err := os.Open(taskFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []models.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(tasks)
}
