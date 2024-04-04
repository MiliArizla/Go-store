package model

import "fmt"

type User struct {
	Id   string
	UserName    string
	Email	string
	senha	string
}

func UserFunc() {
	fmt.Printf("oi")
}