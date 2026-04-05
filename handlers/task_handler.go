package handlers

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/models"
	"github.com/RevDau-PallaviShinde/Go-Cli-App/storage"
)

const (
	statusPending    = "pending"
	statusInProgress = "in-progress"
	statusCompleted  = "completed"
)

const (
	priorityLow    = "low"
	priorityMedium = "medium"
	priorityHigh   = "high"
)

func AddTask(title, description, dueDate, priority string) error {
	if title == "" {
		return errors.New("title is required")
	}

	if dueDate != "" {
		if _, err := time.Parse("2006-01-02", dueDate); err != nil {
			return errors.New("invalid due date format (use YYYY-MM-DD)")
		}
	}

	if priority == "" {
		priority = priorityMedium
	}

	if !isValidPriority(priority) {
		return errors.New("invalid priority (low, medium, high)")
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	task := models.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      statusPending, // default status
		Priority:    priority,      // default priority
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, task)
	return storage.SaveTasks(tasks)
}

func ListTasks(status, priority string) ([]models.Task, error) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return nil, err
	}

	var filtered []models.Task

	for _, task := range tasks {
		if status != "" && task.Status != status {
			continue
		}
		if priority != "" && task.Priority != priority {
			continue
		}
		filtered = append(filtered, task)
	}

	return filtered, nil
}

func UpdateTask(id int, title, description, dueDate, status, priority string) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {

			if title != "" {
				task.Title = title
			}

			if description != "" {
				task.Description = description
			}

			if dueDate != "" {
				if _, err := time.Parse("2006-01-02", dueDate); err != nil {
					return errors.New("invalid due date format (use YYYY-MM-DD)")
				}
				task.DueDate = dueDate
			}

			if status != "" {
				if !isValidStatus(status) {
					return errors.New("invalid status (pending, in-progress, completed)")
				}
				task.Status = status
			}

			if priority != "" {
				if !isValidPriority(priority) {
					return errors.New("invalid priority (low, medium, high)")
				}
				task.Priority = priority
			}

			task.UpdatedAt = time.Now()
			tasks[i] = task

			return storage.SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func DeleteTask(id int) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return storage.SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func SearchTasks(keyword string) ([]models.Task, error) {
	if keyword == "" {
		return nil, errors.New("keyword is required")
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return nil, err
	}

	var results []models.Task
	keyword = strings.ToLower(keyword)

	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), keyword) ||
			strings.Contains(strings.ToLower(task.Description), keyword) {
			results = append(results, task)
		}
	}

	return results, nil
}

func ExportTasksToCSV(fileName string) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"ID", "Title", "Description", "Status", "Priority", "DueDate",
	})

	for _, task := range tasks {
		writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Title,
			task.Description,
			task.Status,
			task.Priority,
			task.DueDate,
		})
	}

	return nil
}

func GetSummary() (int, int, int, error) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return 0, 0, 0, err
	}

	total := len(tasks)
	completed := 0
	pending := 0

	for _, task := range tasks {
		switch task.Status {
		case statusCompleted:
			completed++
		case statusPending:
			pending++
		}
	}

	return total, completed, pending, nil
}

func isValidStatus(status string) bool {
	return status == statusPending ||
		status == statusInProgress ||
		status == statusCompleted
}

func isValidPriority(priority string) bool {
	return priority == priorityLow ||
		priority == priorityMedium ||
		priority == priorityHigh
}
