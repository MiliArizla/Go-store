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

func SearchProductsByNameHandler( w http.ResponseWriter, r *http.Request ) {

}

func CreateProductHandler( w http.ResponseWriter, r *http.Request ) {

}
