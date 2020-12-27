package main

import (
	"log"
    "net/http"
	"github.com/vsergio011/apitasks/routes"

)

func main() {
	s := routes.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
  
}

