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
			w.Header().Add("Content-Type", "application/json")
			//Necessary for request cords
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			w.Header().Set("authorization", "Bearer ACCESS_TOKEN")
			fmt.Println("texto header")
			fmt.Println(r.Header.Get("Authorization"))
			reqToken := r.Header.Get("Authorization")
			//reqToken := "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImNjM2Y0ZThiMmYxZDAyZjBlYTRiMWJkZGU1NWFkZDhiMDhiYzUzODYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibWluYW1lIiwicGljdHVyZSI6Imh0dHA6Ly93d3cuZXhhbXBsZS5jb20vMTIzNDU2NzgvcGhvdG8ucG5nIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL215dGltZS04MjI5MSIsImF1ZCI6Im15dGltZS04MjI5MSIsImF1dGhfdGltZSI6MTYyMDEyMTM0NCwidXNlcl9pZCI6Im5uYUdsa093OXVQSlZycWYxcHdwcHhPU1hPYzIiLCJzdWIiOiJubmFHbGtPdzl1UEpWcnFmMXB3cHB4T1NYT2MyIiwiaWF0IjoxNjIwMTIxMzQ0LCJleHAiOjE2MjAxMjQ5NDQsImVtYWlsIjoibWllbWFpbEBnbWFsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJwaG9uZV9udW1iZXIiOiIrMzQ2NTg0NTIzNjUiLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7InBob25lIjpbIiszNDY1ODQ1MjM2NSJdLCJlbWFpbCI6WyJtaWVtYWlsQGdtYWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.j9R_lJO5zh1WCgPXzpwhpl9kcYtHa81qD4n8s2JxNfTzucLmYPG2fSYUrVa14XxLGqaJE7HvsrwR1hOTnuhUj26PuqbVd3fQQIKi4O6_W1NEFd98Skpe2SS5OXNmbbVuNSJiMv-k7cKI8NtkYxeU68NSX3HIxrOyzbiQ5BkaoWGs1ZkHifT1JEL7dvLjcDMrYk2lGjtPNF_9ea79XwzbpQ3_buwZLv8bR525OmtPDrOqD6odKExYbjOyWUv491fL0J2lujrSK-DgQtWdUZ7vYX-9A_tg9kqLcp5E3EI7jZdH-b0zvb6B_IoGbiUpi2vwTVua0-9qtj1fe0DUmOBm4g"
			if len(reqToken) > 0 {
				splitToken := strings.Split(reqToken, "Bearer ")
				reqToken = splitToken[1]
				fmt.Println("Authorization ENDED")
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
	apiTest.HandleFunc("/testt", a.example).Methods("get")
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
	apiRoutes.HandleFunc("/users", a.fetchUsers).Methods("get", "OPTIONS")
	apiRoutes.HandleFunc("/user/{ID:[a-zA-Z0-9_]+}", a.fetchUser).Methods("get", "OPTIONS")
	apiRoutes.HandleFunc("/adduser", a.addUser).Methods("POST", "OPTIONS")
	apiRoutes.HandleFunc("/deluser/{ID:[a-zA-Z0-9_]+}", a.removeUser).Methods("get", "OPTIONS")
	apiRoutes.HandleFunc("/getToken", a.getToken).Methods("get", "OPTIONS")
	//Tasks routes
	apiRoutes.HandleFunc("/tasks", a.fetchTasks).Methods("get", "OPTIONS")
	apiRoutes.HandleFunc("/task/{id}", a.fetchTask).Methods("get", "OPTIONS")
	apiRoutes.HandleFunc("/taskscreatedby/{id}", a.fetchTasksCreatedBy).Methods("get", "OPTIONS")
	apiRoutes.HandleFunc("/addtask", a.addTask).Methods("POST", "OPTIONS")
	apiRoutes.HandleFunc("/updatetask", a.updateTask).Methods("POST", "OPTIONS")
	apiRoutes.HandleFunc("/deletetask", a.deleteTask).Methods("POST", "OPTIONS")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
