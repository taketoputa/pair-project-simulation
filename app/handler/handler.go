package handler

import (
	"database/sql"
	"log"
	"pair-project/entity"
	"time"
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

func SalesRecap(db *sql.DB, startDate, endDate time.Time) (int, int, float64, error) {
	query := `SELECT COUNT(*) as total_sales , SUM(s.quantity) as total_item_sold , SUM(p.price * s.quantity) as total_revenue from sales s join products p ON s.product_id = p.id where s.sale_date between ? AND ?`
	var totalSales int
	var totalItemsSold int
	var totalRevenue float64

	err := db.QueryRow(query, startDate, endDate).Scan(&totalSales, &totalItemsSold, &totalRevenue)
	if err != nil {
		return 0, 0, 0, err
	}
	return totalSales, totalItemsSold, totalRevenue, nil
}
