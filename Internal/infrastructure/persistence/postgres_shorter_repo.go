package persistence

import (
	"URL_shortener/Internal/domain"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
)

type PostgresShorterRepository struct {
	db *sql.DB
}

func NewPostgresShorterRepository(dataSourceName string) (*PostgresShorterRepository, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresShorterRepository{db: db}, nil
}

func (repo *PostgresShorterRepository) Shorten(url *domain.URL) (string, error) {
	id := uuid.New().String()
	shortURL := generateShortURL(id)

	err := repo.SaveURLMapping(id, url.String(), shortURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (repo *PostgresShorterRepository) SaveURLMapping(id, originalURL, shortURL string) error {
	query := `INSERT INTO url_shortener (id, original_url, short_url) VALUES ($1, $2, $3)`
	_, err := repo.db.Exec(query, id, originalURL, shortURL)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return err
}

func generateShortURL(id string) string {
	return "http://short.url/" + id
}
