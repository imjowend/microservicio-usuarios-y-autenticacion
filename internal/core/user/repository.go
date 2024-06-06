package user

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *Repository) CreateReport(ctx context.Context) error {
	return nil
}
