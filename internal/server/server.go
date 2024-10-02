package server

import (
	"database/sql"
	"fmt"
	"net/http"
)

func StartServer(db *sql.DB, port string) {
	http.HandleFunc("/", HomeHandler(db))

	fmt.Printf("Starting server on port%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
