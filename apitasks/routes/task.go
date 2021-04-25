package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vsergio011/apitasks/controllers"
)

func (a *api) fetchTasks(w http.ResponseWriter, r *http.Request) {
	data, err := controllers.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/tasks]"))
	}
	fmt.Println(data)
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
func (a *api) fetchTask(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idHeader := params["id"]
	fmt.Println("IDHEADER" + idHeader)
	if idHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CANT FIND ID PARAMETER IN QUERY"))
	}

	id, err := strconv.ParseInt(idHeader, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID PARAMTER WITH BAD FORMAT"))
	}

	data, err := controllers.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/task]"))
	}

	json.NewEncoder(w).Encode(data)

}

func (a *api) fetchTasksCreatedBy(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idHeader := params["id"]
	fmt.Println("IDHEADER" + idHeader)
	if idHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CANT FIND ID PARAMETER IN QUERY"))
	}
	//comentado porque funcionaba para id ahora se cambia uid
	/*id, err := strconv.ParseInt(idHeader, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID PARAMTER WITH BAD FORMAT"))
	}*/

	data, err := controllers.GetTasksCreatedBy("prueba")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/task]"))
	}

	json.NewEncoder(w).Encode(data)

}
