package model

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
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

func SearchProductByName(productName string) Product {
	db := db.ConnectToDataBase()

	res := db.QueryRow("SELECT * FROM product where name = $1", productName)

	err := res.Scan(&id, &name, &description, &category, &price, &quantity, &image)
	if err == sql.ErrNoRows {
		fmt.Printf("Product not found %s\n", name)

	} else if err != nil {
		panic(err.Error())
	}

	var p Product
	p.ProductId = id
	p.Name = name
	p.Description = description
	p.Category = category
	p.Price = price
	p.Quantity = quantity
	p.Image = image

	defer db.Close()
	return p
	
}

func CreateProduct(prod Product) error {
	
	if productExist(prod.Name) {
		fmt.Printf("Product already exists: %s\n", prod.Name)
		return fmt.Errorf("product already exists")
	}

	db := db.ConnectToDataBase()
	defer db.Close()
	id := uuid.NewString()
	name := prod.Name
	description := prod.Description
	category := prod.Category
	price := prod.Price
	quantity := prod.Quantity
	image := prod.Image

	strInsert := "INSERT INTO product VALUES($1, $2, $3, $4, $5, $6, $7)"

	result, err := db.Exec(strInsert, id, name, description, category, strconv.FormatFloat(price, 'f', 1, 64), strconv.Itoa(quantity), image)

	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Product %s created (%d row affected)\n", id, rowsAffected)

	return nil

}

func productExist(nameProduct string) bool {
	prod := SearchProductByName(nameProduct)

	return prod.Name == nameProduct
}

func RemoveProduct(id string) error{
	db := db.ConnectToDataBase()
	defer db.Close()

	result, err := db.Exec("DELETE FROM product WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("id: ", id)
	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	fmt.Printf("Product %s deleted (%d row affected)\n", id, rowsAffected)

	return nil
}

func UpdateProduct(prod Product) error {
	db := db.ConnectToDataBase()
	defer db.Close()

	id := prod.ProductId
	name := prod.Name
	description := prod.Description
	// category := prod.Category
	// price := prod.Price
	// quantity := prod.Quantity
	// image := prod.Image

	result, err := db.Exec("UPDATE product SET name= $1, description= $2 where id= $3", name, description, id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	fmt.Printf("Product %s updated (%d row affected)\n", id, rowsAffected)

	return nil
}
