package model

import "fmt"

type Cart struct {
	Id	string
	UserId	string
	InfoProduct []InfoProduct
	Total float64
	// InfoProduct map[string]int
}

type InfoProduct struct {
	ProductId string
	ProductQuantity int
}

func CartFunc() {
	fmt.Printf("oi")
}