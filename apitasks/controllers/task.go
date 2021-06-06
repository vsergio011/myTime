package controllers

import (
	"github.com/vsergio011/apitasks/models"
)

func GetTasks() ([]models.Task, error) {

	return models.GetTasks()
}

func GetTask(id int64) (*models.Task, error) {
	return models.GetTask(id)
}
func GetTasksCreatedBy(id string) ([]models.Task, error) {

	return models.GetTaskscreatedBy(id)
}

func AddTask(task *models.Task) (*models.Task, error) {
	return models.InsertTask(task)
}

func UpdateTask(task *models.Task) (*models.Task, error) {
	return models.UpdateTask(task)
}

func DeleteTask(task *models.Task) (*models.Task, error) {
	return models.DeleteTask(task)
}
