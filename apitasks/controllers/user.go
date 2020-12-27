package controllers
import (
	"github.com/vsergio011/apitasks/models"
)

func GetUsers() ([]models.User, error) {
	
	return models.GetUsers()
}