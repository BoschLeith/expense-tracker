package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/BoschLeith/expense-tracker/internal/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	expenses *sqlite.ExpenseModel
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		expenses: &sqlite.ExpenseModel{
			DB: db,
		},
	}

	server := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	fmt.Printf("Starting server on port%s", server.Addr)
	server.ListenAndServe()
}
