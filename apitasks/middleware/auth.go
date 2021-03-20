package middleware

import (
	"fmt"
	"net/http"
)

func authmiddleware(siguienteManejador http.Handler) http.Handler {
	return http.HandlerFunc(
		func(respuesta http.ResponseWriter, r *http.Request) {
			fmt.Println("hola soy el middle")
			prueba := r.Header.Get("prueba")
			fmt.Println(prueba)
			// Si no llamamos a siguienteManejador, se detiene
			// así que podemos aquí comprobar algo y detener determinada acción
			// Por ejemplo, permitir solo si no son de tipo DELETE
			if r.Method == http.MethodDelete {
				http.Error(respuesta, "Permiso denegado", http.StatusForbidden)
			} else {
				// En caso de que sea permitida llamamos a siguienteManejador
				// y le pasamos la respuesta con la petición
				siguienteManejador.ServeHTTP(respuesta, r)
			}
		})
}
