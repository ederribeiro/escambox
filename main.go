package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

    "github.com/gorilla/mux"
	"github.com/ederribeiro/escambox/src/models"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home")
	fmt.Println("Endpoint Hit: homePage")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{Title: "Camisa", Description: "Camisa azul"}
	fmt.Println("Endpoint Hit: returnProduct")
	json.NewEncoder(w).Encode(product)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/product", getProducts).Methods("GET")
    log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}