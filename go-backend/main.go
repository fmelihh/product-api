package main

import (
	"log"
	"net/http"

	"product-api-go/api/routes"
	_ "product-api-go/docs"
	"product-api-go/internal/db"
	"product-api-go/internal/product"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

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

	repo := product.NewInMemoryRepository()
	service := product.NewService(repo)
	productHandler := product.NewHandler(service)

	routes.RegisterProductRoutes(router, productHandler)

	// database initialization
	err := db.InitDatabase(":memory:")
	if err != nil {
		log.Fatal("There's error with the server", err)
		return
	}

	log.Println("Server Starting on 8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
