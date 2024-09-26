package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.tools.sap/I586129/loja-digport-backend/db"
)

func main() {
	fmt.Println("Server Started")
	StartServer()
	db.InitDB()
}
