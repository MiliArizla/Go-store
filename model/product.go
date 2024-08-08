package model

import (
	"github.tools.sap/I586129/loja-digport-backend/db"
)

type Product struct {
	ProductId   string
	Name        string
	Description string
	Category    string
	Price       float64
	Quantity    int
	Image       string
}

type NewProduct struct {
	Name        string
	Description string
	Category    string
	Price       float64
	Quantity    int
	Image       string
}

var id, name, description, category string
var price float64
var quantity int
var image string

func SearchAllProducts() []Product {

	db := db.ConnectToDataBase()

	result, err := db.Query("SELECT * FROM product")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for result.Next() {
		err = result.Scan(&id, &name, &description, &category, &price, &quantity, &image)

		if err != nil {
			panic(err.Error())
		}

		p.ProductId = id
		p.Name = name
		p.Description = description
		p.Category = category
		p.Price = price
		p.Quantity = quantity
		p.Image = image

		products = append(products, p)
	}

	defer db.Close()

	return products

}