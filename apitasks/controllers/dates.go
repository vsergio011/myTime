package controllers
import (
	"github.com/vsergio011/apitasks/models"
)

func FetchDates() models.Date {
	d := models.Date{Name: "prueba"}
	d.Position = "posicion"
	return d
}