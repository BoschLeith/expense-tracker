package models

import "time"

type Expense struct {
	ID        int
	Name      string
	Category  string
	Amount    float32
	CreatedAt time.Time
}
