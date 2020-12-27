package routes
import (
	"net/http"
	"encoding/json"
	"github.com/vsergio011/apitasks/controllers"
)
func (a *api) fetchTasks(w http.ResponseWriter, r *http.Request) {
	data, err := controllers.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/tasks]"))
	}
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
