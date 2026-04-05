package main

import (
	"log"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found, relying on system environment") // production uses system environment variables
	}

	db, err := database.DbInit()
	if err != nil {
		log.Fatalf("Error: Failed to connect to database\n%v\n", err)
	}

	internal.Run(db)
}
