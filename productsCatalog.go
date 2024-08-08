package main

import (
	"fmt"

	"github.tools.sap/I586129/loja-digport-backend/model"

	"github.com/google/uuid"
)

var productsCatalog []model.Product = []model.Product {
		{
		ProductId: "3849187d-7090-4259-b1f8-8c6116991c4e",
		Name: "Quadrinho Decorativo",
		Description: "Quadrinho Decorativo",
		Category: "Quadro",
		Price: 15.99,
		Quantity: 5,
		Image: "image",
		},
		{
			ProductId: "9d47171a-a730-4cdd-8218-eab1608fa93f",
			Name: "Planta de Teto",
			Description: "Planta de Teto",
			Category: "Planta",
			Price: 13.99,
			Quantity: 10,
			Image: "Image2",
		},
		{
			ProductId: "c972df9d-8d24-4250-8fd9-3757ed06efb2",
			Name: "Poltrona Branca",
			Description: "Poltrona Branca",
			Category: "Poltrona",
			Price: 60.00,
			Quantity: 20,
			Image: "Image3",
		},
		{
			ProductId: "3a0fd158-270a-457e-b9f3-4f23bae361a8",
			Name: "Abajur Moderno",
			Description: "Abajur com design moderno e sofisticado",
			Category: "Iluminação",
			Price: 35.00,
			Quantity: 15,
			Image: "Image4",
	},
	{
			ProductId: "d1692f21-f259-46c1-8e60-15b915c85ea8",
			Name: "Tapete Persa",
			Description: "Tapete Persa, trama em algodão com detalhes em vermelho",
			Category: "Tapetes",
			Price: 120.00,
			Quantity: 10,
			Image: "Image5",
	},
	{
			ProductId: "dd7f015b-9a3a-4f91-9f3d-8e35df602009",
			Name: "Vaso de Cerâmica",
			Description: "Vaso de Cerâmica, cor terra cota, perfeito para plantas de interior",
			Category: "Vasos",
			Price: 18.00,
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


func registerProduct(newProduct model.NewProduct) {
	newId := uuid.New().String()
	product := model.Product {
		ProductId: newId,
		Name: newProduct.Name,
		Description: newProduct.Description,
		Category: newProduct.Category,
		Price: newProduct.Price,
		Quantity: newProduct.Quantity,
		Image: newProduct.Image,
	}
	productsCatalog = append(productsCatalog, product)
}