package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/products", productsHandler)

	http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	
	name := r.URL.Query().Get("name")
	
	if name != "" {
		productName, err := getProductByName(name)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(productName)
		}
	} else {

		products := productsCatalog()

		json.NewEncoder(w).Encode(products)

	}
}