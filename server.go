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
	products := productsCatalog()
	
	json.NewEncoder(w).Encode(products)
}
