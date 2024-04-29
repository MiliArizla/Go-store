package main

import (
	"fmt"

	"github.tools.sap/I586129/loja-digport-backend/model"
)

var productsCatalog []model.Product = []model.Product {
		{
		ProductId: "01",
		Name: "Quadrinho Decorativo",
		Description: "Quadrinho Decorativo",
		Category: "Quadro",
		Value: 15.99,
		Quantity: 5,
		Image: "image",
		},
		{
			ProductId: "02",
			Name: "Planta de Teto",
			Description: "Planta de Teto",
			Category: "Planta",
			Value: 13.99,
			Quantity: 10,
			Image: "Image2",
		},
		{
			ProductId: "03",
			Name: "Poltrona Branca",
			Description: "Poltrona Branca",
			Category: "Poltrona",
			Value: 60.00,
			Quantity: 20,
			Image: "Image3",
		},
		{
			ProductId: "04",
			Name: "Abajur Moderno",
			Description: "Abajur com design moderno e sofisticado",
			Category: "Iluminação",
			Value: 35.00,
			Quantity: 15,
			Image: "Image4",
	},
	{
			ProductId: "05",
			Name: "Tapete Persa",
			Description: "Tapete Persa, trama em algodão com detalhes em vermelho",
			Category: "Tapetes",
			Value: 120.00,
			Quantity: 10,
			Image: "Image5",
	},
	{
			ProductId: "06",
			Name: "Vaso de Cerâmica",
			Description: "Vaso de Cerâmica, cor terra cota, perfeito para plantas de interior",
			Category: "Vasos",
			Value: 18.00,
			Quantity: 30,
			Image: "Image6",
	},
};


func getProductByName(name string) (*model.Product, error) {
	for _, product := range productsCatalog {
		if product.Name == name {
			return &product, nil
	}
}
	return nil, fmt.Errorf("product does not exist")
};

func registerProduct(newProduct model.Product) error {
	for _, products := range productsCatalog{
		if products.ProductId == newProduct.ProductId {
			return fmt.Errorf("ID Repetido")
		}
	}
	productsCatalog = append(productsCatalog, newProduct)
	return nil
}