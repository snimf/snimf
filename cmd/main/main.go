package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sn1mf/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
