package model

type Product struct {
	ProductId   string
	Name        string
	Description string
	Category    string
	Value       float64
	Quantity    int
	Image       string
}

type NewProduct struct {
	Name        string
	Description string
	Category    string
	Value       float64
	Quantity    int
	Image       string
}
