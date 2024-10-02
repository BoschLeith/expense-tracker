package server

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/BoschLeith/expense-tracker/internal/database"
)

// TODO: Remove database logic from handler
func HomeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenses, err := database.GetExpenses(db)
		if err != nil {
			http.Error(w, "Unable to query expenses", http.StatusInternalServerError)
			return
		}

		tmplPath := filepath.Join("cmd", "web", "index.html")
		tmpl := template.Must(template.ParseFiles(tmplPath))
		tmpl.Execute(w, expenses)
	}
}
