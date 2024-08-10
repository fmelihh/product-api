package main

import (
	"fmt"
	"log"
	"net/http"

	_ "product-api-go/api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Product struct {
	name        string  `json:"name"`
	price       float32 `json:"price"`
	description string  `json:"description"`
	currency    string  `json:"currency"`
	category    string  `json:"category"`
	location    string  `json:"location"`
}

// @Summary Get Product
// @Description get product
// @Produce  json
// @Success 200
// @Router /product/{product-id} [get]
func GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("get product")
}

// @Summary Insert Product
// @Description insert product
// @Produce  json
// @Success 200
// @Router /product [post]
func AddProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("post product")
}

// @Summary List Products
// @Description list inserted products
// @Produce  json
// @Success 200 {array} Product
// @Router /product/list [get]
func ListProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("list product")
}

// @Summary Delete Product
// @Description delete inserted product
// @Produce  json
// @Success 200
// @Router /product/ [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("delete product")
}

// @Summary Update Product
// @Description update inserted product
// @Produce  json
// @Success 200
// @Router /product/ [put]
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Print("update product")
}

// @title Product API
// @version 1.0
// @description CRUD Operations for the products
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
func main() {
	// Swagger URL: http://localhost:8000/docs/index.html
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	productRouter := router.PathPrefix("/products/").Subrouter()
	productRouter.HandleFunc("", AddProduct).Methods("POST")
	productRouter.HandleFunc("", UpdateProduct).Methods("PUT")
	productRouter.HandleFunc("/list", ListProduct).Methods("GET")
	productRouter.HandleFunc("", DeleteProduct).Methods("DELETE")
	productRouter.HandleFunc("/{productID}", GetProduct).Methods("GET")

	log.Println("Server Starting on 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
