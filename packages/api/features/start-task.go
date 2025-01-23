package features

import (
	"task-manager/packages/database"
	"task-manager/packages/database/models"
)

func StartTask(id string) {
	db := database.Connect()

	var task models.Task

	db.First(&task, "id = ?", id)

	task.Status = models.InProgress

	db.Save(&task)
}
