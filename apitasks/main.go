package main

import (
    //"context"
	"log"
	//"os"
    //"fmt"
    "net/http"
  /*  "os/signal"
	"syscall"
	"time"*/

   /*"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"*/
	"github.com/vsergio011/apitasks/routes"
	//"github.com/urfave/cli/v2"
    //"go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	s := routes.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
  
}

