package main

import (
	"log"
	"net/http"

	"github.com/Karthik-S-L/library-management-golan/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe("localhost:4000",router))

}