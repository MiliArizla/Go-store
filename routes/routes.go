package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.tools.sap/I586129/loja-digport-backend/controller"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/product", controller.SearchProductsHandler).Methods("GET")
	route.HandleFunc("/product-name", controller.SearchProductByNameHandler).Methods("GET")
	route.HandleFunc("/products", controller.CreateProductHandler).Methods("POST")

	http.ListenAndServe(":8080", route)
}