package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ankush-web-eng/Bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
}

func main() {
	fmt.Println("Server started at 0.0.0.0:9010")
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", nil))
}
