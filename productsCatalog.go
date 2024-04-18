package main

import (
	"github.tools.sap/I586129/loja-digport-backend/model"
)

func productsCatalog() []model.Product {
	productsCatalog := []model.Product {
		{
		ProductId: "01",
		Name: "Quadro ",
		Description: "Blusa Rosa Feminina manga longa",
		Category: "Blusa",
		Value: 25.99,
		Quantity: 5,
		Image: "image",
		},
		{
			ProductId: "02",
			Name: "Blusa Verde",
			Description: "Blusa verde feminina manga curta",
			Category: "Blusa",
			Value: 26.99,
			Quantity: 10,
			Image: "Image2",
		},
		{
			ProductId: "03",
			Name: "Camisa",
			Description: "Camisa branca masculina",
			Category: "Camisa",
			Value: 30.00,
			Quantity: 20,
			Image: "https://todaatual.com/wp-content/uploads/2015/02/QuadrosDecorativosParaSala-1.jpg",
		},
}
	return productsCatalog
};