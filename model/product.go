package model

import "fmt"

type Product struct {
	ProductId    string
	Name   string
	Description string
	Category   string
	Value float64
	Quantity int
	Image string
}

func ProductFunc() {
	fmt.Printf("oi")
}