package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.tools.sap/I586129/loja-digport-backend/model"
)

func SearchProductsHandler( w http.ResponseWriter, r *http.Request ) {
	product := model.SearchAllProducts()
	json.NewEncoder(w).Encode(product)
}

func SearchProductByNameHandler( w http.ResponseWriter, r *http.Request ) {
	name := r.URL.Query().Get("name")
	product := model.SearchProductByName(name)
	json.NewEncoder(w).Encode(product)
}

func CreateProductHandler( w http.ResponseWriter, r *http.Request ) {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	error := model.CreateProduct(product)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveProductHandler(w http.ResponseWriter, r *http.Request) {
	// implementation of the RemoveProdutoHandler function
	// the function should receive a request and remove a product from the database
	// the product to be removed should be passed as a parameter in the request body
	// the function should return a status code 204 if the product was removed successfully, no content
	// or a status code 404 if the product was not found
	id := mux.Vars(r)["id"]
	error := model.RemoveProduct(id)
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// implementation of AtualizaProdutoHandler
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)
	error := model.UpdateProduct(product)
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Println(product.ProductId)
		w.WriteHeader(http.StatusOK)
	}

}
