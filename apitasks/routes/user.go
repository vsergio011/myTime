package routes
import (
	"net/http"
	"encoding/json"
	"github.com/vsergio011/apitasks/controllers"
)
func (a *api) fetchUsers(w http.ResponseWriter, r *http.Request) {
	data, err := controllers.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/user]"))
	}
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}


