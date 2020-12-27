package routes

import (
"github.com/gorilla/mux"
"net/http"
"encoding/json"
"github.com/vsergio011/apitasks/controllers"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/tasks",a.fetchTasks).Methods("get")
	

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
	
type person struct {
    name string
    age  int
}
func (a *api) fetchTasks(w http.ResponseWriter, r *http.Request) {
	data, err := controllers.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/user]"))
	}
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
