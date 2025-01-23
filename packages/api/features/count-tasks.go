package features

import (
	"task-manager/packages/database"
	"task-manager/packages/database/models"
)

type CountTasksInput struct {
	Status string
}

func CountTasks(input CountTasksInput) int64 {
	db := database.Connect()

	var count int64

	if input.Status != "" {
		db.Model(&models.Task{}).Count(&count).Where("status = ?", input.Status)
	} else {
		db.Model(&models.Task{}).Count(&count)
	}

	return count
}
