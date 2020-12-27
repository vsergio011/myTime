package controllers
import (
	"github.com/vsergio011/apitasks/models"
)

func GetTasks() ([]models.Task, error) {
	
	return models.GetTasks()
}