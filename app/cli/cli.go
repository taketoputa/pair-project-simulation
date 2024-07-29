package cli

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
	"pair-project/handler"
	"strconv"
	"time"
)

func AddProductCLI(db *sql.DB) {
	var name string
	var price float64
	var stock int

	fmt.Println("Enter product name:")
	fmt.Scanln(&name)

	// validation for price input
	for {
		fmt.Println("Enter product price:")
		var priceInput string
		fmt.Scanln(&priceInput)
		if p, err := strconv.ParseFloat(priceInput, 64); err == nil {
			price = p
			break
		} else {
			fmt.Println("Invalid input. Please enter a numeric value for price.")
		}
	}

	// validation for stock input
	for {
		fmt.Println("Enter product stock:")
		var stockInput string
		fmt.Scanln(&stockInput)
		if s, err := strconv.Atoi(stockInput); err == nil {
			stock = s
			break
		} else {
			fmt.Println("Invalid input. Please enter a numeric value for stock.")
		}
	}

	product := entity.Product{Name: name, Price: price, Stock: stock}
	handler.AddProduct(db, product)
}

func UpdateStockCLI(db *sql.DB) {
	var id, stockChange int

	// validation for product ID input
	for {
		fmt.Println("Enter product ID to update stock:")
		var idInput string
		fmt.Scanln(&idInput)
		if i, err := strconv.Atoi(idInput); err == nil {
			id = i
			// cek ID, exists or not
			if handler.ProductExists(db, id) {
				break
			} else {
				fmt.Println("Product ID does not exist. Please enter a valid product ID.")
			}
		} else {
			fmt.Println("Invalid input. Please enter a numeric value for product ID.")
		}
	}

	// validation for check input
	for {
		fmt.Println("Enter stock change amount (+ to add, - to reduce):")
		var stockChangeInput string
		fmt.Scanln(&stockChangeInput)
		if sc, err := strconv.Atoi(stockChangeInput); err == nil {
			stockChange = sc
			break
		} else {
			fmt.Println("Invalid input. Please enter a numeric value for stock change.")
		}
	}

	handler.UpdateStock(db, id, stockChange)
}

/*
*
Temuan:
- Clean

Saran:
*/
func AddStaffCLI(db *sql.DB) {
	var name, email, position string

	fmt.Println("Enter staff name:")
	fmt.Scanln(&name)
	fmt.Println("Enter staff email:")
	fmt.Scanln(&email)
	fmt.Println("Enter staff position:")
	fmt.Scanln(&position)

	staff := entity.Staff{Name: name, Email: email, Position: position}
	handler.AddStaff(db, staff)
}

/*
*
Temuan:
Fungsi Sales Recap Fatal Error dan keluar aplikasi jika tidak ada data ditemukan

Saran:
Print "No data found for the specified date range."
*/
func SalesRecapCLI(db *sql.DB) {
	var startDateStr, endDateStr string
	fmt.Println("Enter start date (YYYY-MM-DD):")
	fmt.Scanln(&startDateStr)
	fmt.Println("Enter end date (YYYY-MM-DD):")
	fmt.Scanln(&endDateStr)

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		log.Fatalf("Invalid start date : %v", err)
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		log.Fatalf("Invalid end date : %v", err)
	}

	totalSales, totalItemsSold, totalRevenue, err := handler.SalesRecap(db, startDate, endDate)
	if err != nil {
		fmt.Println("no data found for the specified date range.")
		return
	}

	//display the result
	fmt.Printf("Sales Recap from %s to %s\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	fmt.Printf("Total Sales : %d\n", totalSales)
	fmt.Printf("Total items Sold : %d\n", totalItemsSold)
	fmt.Printf("Total Revenue : %.2f\n", totalRevenue)
}

func RunCLI(db *sql.DB) {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Product Stock")
		fmt.Println("3. Add Staff")
		fmt.Println("4. Sales Recap")
		fmt.Println("5. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddProductCLI(db)
		case 2:
			UpdateStockCLI(db)
		case 3:
			AddStaffCLI(db)
		case 4:
			SalesRecapCLI(db)
		case 5:
			fmt.Println("Exit")
			return
		}
	}
}
