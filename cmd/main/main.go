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

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func main() {
	fmt.Println("Server started at localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
