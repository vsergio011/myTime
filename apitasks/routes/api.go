package routes

import (
"github.com/gorilla/mux"
"net/http"
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
	//Users routes
	//r.HandleFunc("/users/add", signInController).Methods("POST")
	r.HandleFunc("/users",a.fetchUsers).Methods("get")

	//Tasks routes
	r.HandleFunc("/tasks",a.fetchTasks).Methods("get")
	r.HandleFunc("/task/{id}",a.fetchTask).Methods("get")
	r.HandleFunc("/taskscreatedby/{id}",a.fetchTasksCreatedBy).Methods("get")
	

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
	
