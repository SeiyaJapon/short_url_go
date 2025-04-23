package config

import (
	"URL_shortener/Internal/infrastructure/persistence"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitializeRepositories() (*persistence.PostgresShorterRepository, *persistence.PostgresRedirectRepository, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dataSourceName := os.Getenv("POSTGRES_DB")
	if dataSourceName == "" {
		log.Fatal("POSTGRES_DB environment variable is not set")
	}

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, nil, err
	}

	urlShorterRepository, err := persistence.NewPostgresShorterRepository(dataSourceName)
	if err != nil {
		return nil, nil, err
	}

	urlRedirecterRepository := persistence.NewPostgresRedirectRepository(db)

	return urlShorterRepository, urlRedirecterRepository, nil
}
