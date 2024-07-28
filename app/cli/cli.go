package cli

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
	"pair-project/handler"
	"time"
)

func AddProductCLI(db *sql.DB) {
	var name string
	var price float64
	var stock int

	/**
	- Temuan:
	Price dan stock apabila dimasukkan input non angka, mengakibatkan insert ke database dengan value 0

	- Saran:
	Jika price dan stock yang diinput adalah non angka,
	maka dilooping agar user input angka

	bisa menggunakan loop dengan nama, sehingga breaking loop lebih mudah
	example:
	loop1:
	for true{
		if condition1 {
		break loop1 }
		else { do nothing}
	}
	*/

	fmt.Println("Enter product name:")
	fmt.Scanln(&name)
	fmt.Println("Enter product price:")
	fmt.Scanln(&price)
	fmt.Println("Enter product stock:")
	fmt.Scanln(&stock)

	product := entity.Product{Name: name, Price: price, Stock: stock}
	handler.AddProduct(db, product)
}

/*
*
Temuan:
  - Apabila user input product ID yang tidak ada di database,
    secara program tetap melanjutkan, dan tidak ada logging error
  - stock change apabila dimasukkan input non angka, mengakibatkan insert ke database dengan value 0

Saran:
  - dibuat validasi, product ID yang diinput memang ada di database,
    Jika tidak ada bisa dikembalikan ke Menu utama
  - Sama seperti sebelumnya bisa dibuat validasi
*/
func UpdateStockCLI(db *sql.DB) {
	var id, stockChange int
	fmt.Println("Enter product ID to update stock:")
	fmt.Scanln(&id)
	fmt.Println("Enter stock change amount (+ to add, - to reduce):")
	fmt.Scanln(&stockChange)

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
		log.Fatalf("No data found for the specified date range: %v", err)
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
