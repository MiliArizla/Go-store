package model

import "fmt"

type Wishlist struct {
	Id	string
	ProductId    []string
	UserId   string
}

func WishlistFunc() {
	fmt.Printf("oi")
}