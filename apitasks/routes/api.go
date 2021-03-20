package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vsergio011/apitasks/controllers"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func authmiddleware(siguienteManejador http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hola soy el middle")
			prueba := r.Header.Get("prueba")
			fmt.Println(prueba)
			fmt.Println(r.Header.Get("Authorization"))
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
			data, err := controllers.VerifyToken(reqToken, r.Context())
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Something bad happened! [/tasks]"))
			}

			fmt.Println(data)
			// Si no llamamos a siguienteManejador, se detiene
			// así que podemos aquí comprobar algo y detener determinada acción
			// Por ejemplo, permitir solo si no son de tipo DELETE
			if r.Method == http.MethodDelete {
				http.Error(w, "Permiso denegado", http.StatusForbidden)
			} else {
				// En caso de que sea permitida llamamos a siguienteManejador
				// y le pasamos la respuesta con la petición
				siguienteManejador.ServeHTTP(w, r)
			}
		})
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	//Users routes
	//r.HandleFunc("/users/add", signInController).Methods("POST")
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.Use(authmiddleware)
	apiRoutes.HandleFunc("/users", a.fetchUsers).Methods("get")
	apiRoutes.HandleFunc("/user/{ID:[a-zA-Z0-9_]+}", a.fetchUsers).Methods("get")
	apiRoutes.HandleFunc("/adduser", a.addUser).Methods("POST")
	apiRoutes.HandleFunc("/getToken", a.getToken).Methods("get")
	//Tasks routes
	r.HandleFunc("/tasks", a.fetchTasks).Methods("get")
	r.HandleFunc("/task/{id}", a.fetchTask).Methods("get")
	r.HandleFunc("/taskscreatedby/{id}", a.fetchTasksCreatedBy).Methods("get")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
