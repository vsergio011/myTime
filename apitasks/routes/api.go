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
			fmt.Println(r.Header.Get("Authorization"))
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
			data, err := controllers.VerifyToken(reqToken, r.Context())
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Something bad happened! [/tasks]"))
			}
			//añado el uid del usuario que realiza la peticion
			vars := mux.Vars(r)
			vars["currentUser"] = data.UID

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
	//rutas de test
	apiTest := r.PathPrefix("/test").Subrouter()
	apiTest.HandleFunc("/users", a.fetchUsers).Methods("get")
	apiTest.HandleFunc("/user/{ID:[a-zA-Z0-9_]+}", a.fetchUser).Methods("get")
	apiTest.HandleFunc("/rmuser/{ID:[a-zA-Z0-9_]+}", a.removeUser).Methods("get")
	apiTest.HandleFunc("/upduser/{ID:[a-zA-Z0-9_]+}", a.updateUser).Methods("POST")
	apiTest.HandleFunc("/adduser", a.addUser).Methods("POST")

	//Tasks routes
	apiTest.HandleFunc("/tasks", a.fetchTasks).Methods("get")
	apiTest.HandleFunc("/task/{id}", a.fetchTask).Methods("get")
	apiTest.HandleFunc("/taskscreatedby/{id}", a.fetchTasksCreatedBy).Methods("get")

	//rutas api
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.Use(authmiddleware)
	apiRoutes.HandleFunc("/users", a.fetchUsers).Methods("get")
	apiRoutes.HandleFunc("/user/{ID:[a-zA-Z0-9_]+}", a.fetchUser).Methods("get")
	apiRoutes.HandleFunc("/adduser", a.addUser).Methods("POST")
	apiRoutes.HandleFunc("/deluser/{ID:[a-zA-Z0-9_]+}", a.removeUser).Methods("get")
	apiRoutes.HandleFunc("/getToken", a.getToken).Methods("get")
	//Tasks routes
	apiRoutes.HandleFunc("/tasks", a.fetchTasks).Methods("get")
	apiRoutes.HandleFunc("/task/{id}", a.fetchTask).Methods("get")
	apiRoutes.HandleFunc("/taskscreatedby/{id}", a.fetchTasksCreatedBy).Methods("get")
	apiRoutes.HandleFunc("/addtask", a.fetchTasksCreatedBy).Methods("POST")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
