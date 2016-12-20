package main

import (
	"log"
	"net/http"
	"github.com/ederribeiro/escambox/app/bundles/productsbundle"
	"github.com/gorilla/mux"
)

func main() {
	// Controllers declaration
	pc := &productsbundle.ProductController{}
	r := mux.NewRouter()
 	s := r.PathPrefix("/api/v1/").Subrouter()
	// Routes handling
 	s.HandleFunc("/products", pc.Index).Methods("GET")
	http.Handle("/", r)
 	log.Fatal(http.ListenAndServe(":4040", nil))
}