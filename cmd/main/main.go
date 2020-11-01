package main

import (
	"log"
	"net/http"

	"../../pkg/routes"
	"github.com/gorilla/mux"
	// TODO #3 Get mysql driver
	"gorm.io/driver/mysql"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
