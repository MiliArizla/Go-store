package controller

import (
	"encoding/json"
	"net/http"

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
