package features

import (
	"task-manager/packages/database"
	"task-manager/packages/database/models"
)

type ListTasksInput struct {
	Status models.Status
}

func ListTasks(input ListTasksInput) []models.Task {
	db := database.Connect()

	var tasks []models.Task

	if input.Status != "" {
		db.Find(&tasks, "status = ?", input.Status)
	} else {
		db.Find(&tasks)
	}

	return tasks
}
