-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		amount REAL NOT NULL,
        createdAt DATETIME NOT NULL
	);

INSERT INTO expenses (name, category, amount, createdAt) VALUES ('Groceries', 'Food', 150.75, '2023-10-01 10:30:00');
INSERT INTO expenses (name, category, amount, createdAt) VALUES ('Electricity Bill', 'Utilities', 75.50, '2023-10-02 14:00:00');
INSERT INTO expenses (name, category, amount, createdAt) VALUES ('Car Insurance', 'Transportation', 120.00, '2023-10-05 12:00:00');
INSERT INTO expenses (name, category, amount, createdAt) VALUES ('Internet Bill', 'Utilities', 50.00, '2023-10-07 11:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE expenses;
-- +goose StatementEnd
