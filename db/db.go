package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDataBase() *sql.DB {
	dbPass := os.Getenv("DB_PASS")
	connStr := fmt.Sprint("user=postgres dbname=digport_loja password=", dbPass," host=localhost sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitDB() {
	createTables()
}

func createTables() {
	db := ConnectToDataBase()
	defer db.Close()

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS USUARIO (
		id SERIAL PRIMARY KEY,
		name VARCHAR,
		phone VARCHAR,
		address VARCHAR,
		email VARCHAR NOT NULL UNIQUE,
		password VARCHAR NOT NULL
	);`
	_, err := db.Exec(createUsersTable)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("User Table Created")
	}

	createProductsTableScript := `
	CREATE TABLE IF NOT EXISTS product
	(
    id varchar primary key,
    name varchar,
		description varchar
		category varchar
		price  float8,
		quantity int
		image varchar
	);`

	_, err = db.Exec(createProductsTableScript)
	if err != nil {
		panic("Error creating Product Table")
	} else {
		fmt.Println("Table Product created")
	}
}