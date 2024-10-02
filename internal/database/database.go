package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Expense struct {
	Name     string
	Category string
	Amount   float64
}

func SetupDatabase() (*sql.DB, error) {
	db, err := New()
	if err != nil {
		return nil, err
	}

	err = CreateExpensesTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func New() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./internal/database/test.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateExpensesTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		amount REAL NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func GetExpenses(db *sql.DB) ([]Expense, error) {
	rows, err := db.Query("SELECT name, category, amount FROM expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense
		if err := rows.Scan(&expense.Name, &expense.Category, &expense.Amount); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}
