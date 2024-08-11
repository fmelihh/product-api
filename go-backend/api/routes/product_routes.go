package routes

import (
	"net/http"
	"product-api-go/internal/product"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router, handler *product.Handler) {
	router.HandleFunc("/products", handler.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products", handler.ListProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", handler.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/products/{id}", handler.DeleteProduct).Methods(http.MethodDelete)
}
