package sqlite

import (
	"database/sql"

	"github.com/BoschLeith/expense-tracker/internal/models"
)

type ExpenseModel struct {
	DB *sql.DB
}

func (m *ExpenseModel) GetExpenses() ([]models.Expense, error) {
	stmt := `SELECT id, name, category, amount, createdAt FROM expenses ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	expenses := []models.Expense{}
	for rows.Next() {
		e := models.Expense{}
		err := rows.Scan(&e.ID, &e.Name, &e.Category, &e.Amount, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
