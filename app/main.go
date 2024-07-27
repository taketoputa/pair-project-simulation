package main

import (
	"database/sql"
	"log"
	"pair-project/cli"
	"pair-project/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Open database connection
	db, err := sql.Open("mysql", config.DatabaseConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cli.RunCLI(db)
}
