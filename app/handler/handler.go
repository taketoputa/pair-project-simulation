package handler

import (
	"database/sql"
	"log"
	"pair-project/entity"
)

func AddProduct(db *sql.DB, product entity.Product) {
	query := `INSERT INTO products (name, price, stock) VALUES (?, ?, ?)`
	_, err := db.Exec(query, product.Name, product.Price, product.Stock)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Product added successfully")
}
