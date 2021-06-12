package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vsergio011/apitasks/routes"
)

const version string = "1.0.0"

func main() {

	fmt.Println("He iniciado en el host" + os.Getenv("DB_HOST"))
	s := routes.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}
