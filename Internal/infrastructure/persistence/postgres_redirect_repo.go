package persistence

import (
	"URL_shortener/Internal/domain"
	"database/sql"
)

type PostgresRedirectRepository struct {
	db *sql.DB
}

func NewPostgresRedirectRepository(db *sql.DB) *PostgresRedirectRepository {
	return &PostgresRedirectRepository{db: db}
}

func (repo *PostgresRedirectRepository) Redirect(url *domain.URL) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM url_shortener WHERE short_url = $1`
	err := repo.db.QueryRow(query, url.String()).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
