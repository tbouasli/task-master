package features

import (
	"task-manager/packages/database"
	"task-manager/packages/database/models"
)

type CreateTaskInput struct {
	Name        string
	Description string
}

func CreateTask(input CreateTaskInput) {
	db := database.Connect()

	task := models.NewTask(input.Name, input.Description)

	db.Create(&task)
}
