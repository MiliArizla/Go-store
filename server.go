package main

import (
	"encoding/json"
	"net/http"

	"github.tools.sap/I586129/loja-digport-backend/model"
)

func StartServer() {
	http.HandleFunc("/products", productsHandler)

	http.ListenAndServe(":8080", nil)

}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		getProducts (w,r)
	} else if r.Method == "POST" {
		addProduct (w,r)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	
	if name != "" {
		productName, err := getProductByName(name)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(productName)
		}
	} else {

		products := productsCatalog

		json.NewEncoder(w).Encode(products)

	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct model.NewProduct

	json.NewDecoder(r.Body).Decode(&newProduct)
	
	registerProduct(newProduct)
	w.WriteHeader(http.StatusCreated)

}
