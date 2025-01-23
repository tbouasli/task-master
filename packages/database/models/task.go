package models

import (
	"github.com/google/uuid"
)

type Status string

const (
	NotStarted Status = "Not Started"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
)

type Task struct {
	ID          string
	Name        string
	Description string
	Status      Status
}

func NewTask(name, description string) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Status:      NotStarted,
	}
}
