package config

import (
	"log"
	"os"
)

// Config holds the app configuration.
type Config struct {
	DatabasePath	string
	AdminUsername	string
	AdminPassword	string
}

// LoadConfig loads and returns configuration from env variables.
func LoadConfig() *Config {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./database.db"	// Default sqlite path
	}

	adminUser := os.Getenv("ADMIN_USER")
	adminPass := os.Getenv("ADMIN_PASS")

	if adminUser == "" || adminPass == "" {
		log.Fatal("ADMIN_USER and ADMIN_PASS environment variables must be set.")
	}

	return &Config{
		DatabasePath: dbPath,
		AdminUsername: adminUser,
		AdminPassword: adminPass,
	}
}