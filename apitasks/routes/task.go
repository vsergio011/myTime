package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vsergio011/apitasks/controllers"
	"github.com/vsergio011/apitasks/models"
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

	data, err := controllers.GetTasksCreatedBy(idHeader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/task]"))
	}

	json.NewEncoder(w).Encode(data)

}

func (a *api) addTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entro en el primero")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	var p *models.Task
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := controllers.AddTask(p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func (a *api) updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entro en el primero")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	var p *models.Task
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := controllers.UpdateTask(p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func (a *api) deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entro en el primero")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	var p *models.Task
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := controllers.DeleteTask(p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}
