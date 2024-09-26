package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.tools.sap/I586129/loja-digport-backend/controller"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/products", controller.SearchProductsHandler).Methods("GET")
	route.HandleFunc("/product-name", controller.SearchProductByNameHandler).Methods("GET")
	route.HandleFunc("/add-product", controller.CreateProductHandler).Methods("POST")
	route.HandleFunc("/delete-product/{id}", controller.RemoveProductHandler).Methods("DELETE")
	route.HandleFunc("/product", controller.UpdateProductHandler).Methods("PUT")

	// User
	// route.HandleFunc("/user", controller.CreateUserHandler).Methods("POST")
	// route.HandleFunc("user",controller.SearchUserByEmail).Methods("GET")
	// route.HandleFunc("user/login", controller.LoginHandler).Methods("POST")

	http.ListenAndServe(":8080", route)
}