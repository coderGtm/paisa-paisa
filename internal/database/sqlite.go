package database

import (
	"database/sql"
	"log"

	// Import the sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// DB is the database connection pool (singleton).
var DB *sql.DB

// InitDB initializes the database connection for the app.
func InitDB(databasePath string) {
	var err error
	DB, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established.")
	runMigrations()
}

// runMigrations applies necessary schema changes.
func runMigrations() {
	// TODO: Implement actual migration logic using SQL files in the /migrations directory.
	// For now, we create tables if they don't exist with embedded SQL strings.

	createUsersTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCEMENT,
			username TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			display_name TEXT NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	createCategoriesTableSQL := `
	CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCEMENT,
			parent_id INTEGER,
			name TEXT NOT NULL,
			description TEXT
	);`

	createExpensesTablesSQL := `
	CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCEMENT,
			user_id INTEGER NOT NULL,
			transaction_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			amount FLOAT NOT NULL,
			category_id INTEGER NOT NULL,
			description TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	createSettingsTableSQL := `
	CREATE TABLE IF NOT EXISTS settings (
			id INTEGER PRIMARY KEY AUTOINCEMENT,
			key TEXT NOT NULL,
			value TEXT NOT NULL,
			for_user_id INTEGER NOT NULL,
			updated_by_user_id INTEGER NOT NULL,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(createUsersTableSQL); err != nil {
		log.Fatal("Failed to create users table:", err)
	}
	if _, err := DB.Exec(createCategoriesTableSQL); err != nil {
		log.Fatal("Failed to create categories table:", err)
	}
	if _, err := DB.Exec(createExpensesTablesSQL); err != nil {
		log.Fatal("Failed to create expenses table:", err)
	}
	if _, err := DB.Exec(createSettingsTableSQL); err != nil {
		log.Fatal("Failed to create settings table:", err)
	}
}