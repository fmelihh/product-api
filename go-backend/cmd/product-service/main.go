package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	name        string  `json:"name"`
	price       float32 `json:"price"`
	description string  `json:"description"`
	currency    string  `json:"currency"`
	category    string  `json:"category"`
	location    string  `json:"location"`
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("get product")
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("post product")
}

func listProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("list product")
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("delete product")
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("update product")
}

func main() {
	router := mux.NewRouter()

	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", addProduct).Methods("POST")
	productRouter.HandleFunc("", updateProduct).Methods("PUT")
	productRouter.HandleFunc("/list", listProduct).Methods("GET")
	productRouter.HandleFunc("", deleteProduct).Methods("DELETE")
	productRouter.HandleFunc("/{productID}", getProduct).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
