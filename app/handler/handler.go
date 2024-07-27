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

func UpdateStock(db *sql.DB, id, stockChange int) {
	query := `UPDATE products SET stock = stock + ? WHERE id = ?`
	_, err := db.Exec(query, stockChange, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Stock updated successfully")
}

func AddStaff(db *sql.DB, staff entity.Staff) {
	query := `INSERT INTO staff (name, email, position) VALUES (?, ?, ?)`
	_, err := db.Exec(query, staff.Name, staff.Email, staff.Position)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Staff added successfully")
}
