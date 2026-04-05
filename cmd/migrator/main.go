package main

import (
	"fmt"
	"log"

	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/auth"
	"github.com/Ahmed-Armaan/finance-dashboard-api.git/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found, relying on system environment") // production uses system environment variables
	}

	password := ""
	fmt.Println("Set a password for the default Super admin: super")
	fmt.Scanln(&password)

	db, err := database.DbInit()
	if err != nil {
		log.Fatalf("Error: Migration failed\n%v\n", err)
	}

	hashedPass, err := auth.HashAndSalt([]byte(password))
	if err != nil {
		log.Fatalf("Error: Migration failed\n%v\n", err)
	}

	err = db.Seed(hashedPass)
	if err != nil {
		log.Fatalf("Error: Migration failed\n%v\n", err)
	}
}
