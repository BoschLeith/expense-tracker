package main

import (
	"log"

	"github.com/BoschLeith/expense-tracker/internal/database"
	"github.com/BoschLeith/expense-tracker/internal/server"
)

func main() {
	// TODO: Move SetupDatabase to StartServer
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatalf("Error during database setup: %v", err)
	}
	defer db.Close()

	server.StartServer(db, ":8000")
}
