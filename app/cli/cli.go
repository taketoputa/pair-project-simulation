package cli

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
	"pair-project/handler"
)

func AddProductCLI(db *sql.DB) {
	var name string
	var price float64
	var stock int

	fmt.Println("Enter product name:")
	fmt.Scanln(&name)
	fmt.Println("Enter product price:")
	fmt.Scanln(&price)
	fmt.Println("Enter product stock:")
	fmt.Scanln(&stock)

	product := entity.Product{Name: name, Price: price, Stock: stock}
	handler.AddProduct(db, product)
}

func RunCLI(db *sql.DB) {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Product")
		fmt.Println("5. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddProductCLI(db)
		case 2:
			fmt.Println("exit")
			return
		}
	}
}
