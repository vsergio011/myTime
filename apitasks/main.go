package main

import (
	"log"
	"net/http"

	"github.com/vsergio011/apitasks/routes"
)

const version string = "1.0.0"

func main() {
	s := routes.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}
