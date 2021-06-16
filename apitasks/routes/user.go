package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vsergio011/apitasks/controllers"
	"github.com/vsergio011/apitasks/models"
)

func (a *api) fetchUsers(w http.ResponseWriter, r *http.Request) {
	data, err := controllers.GetUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/user]"))
	}
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("hello world")

	json.NewEncoder(w).Encode(data)
}

func (a *api) fetchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("vars")
	fmt.Println(vars["currentUser"])
	data, err := controllers.GetUser(r.Context(), vars["ID"])
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

func (a *api) addUser(w http.ResponseWriter, r *http.Request) {
	var p models.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := controllers.AddUser(r.Context(), p)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func (a *api) updateUser(w http.ResponseWriter, r *http.Request) {
	var p models.User
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := controllers.UpdateUser(r.Context(), p, vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (a *api) removeUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	vars := mux.Vars(r)
	err := controllers.RemoveUser(r.Context(), vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/user]"))
	}
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func (a *api) getToken(w http.ResponseWriter, r *http.Request) {
	data, err := controllers.CreateToken(r.Context())
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

func (a *api) example(w http.ResponseWriter, r *http.Request) {
	//date, err := controllers.date()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("prueba master")
}
