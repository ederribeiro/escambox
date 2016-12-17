package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"./src/models"
)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home")
	fmt.Println("Endpoint Hit: homePage")
}

func returnProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{Title: "Camisa", Description: "Camisa azul"}
	fmt.Println("Endpoint Hit: returnProduct")
	json.NewEncoder(w).Encode(product)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/product", returnProduct)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}