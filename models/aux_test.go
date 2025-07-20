package models_test

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func emptyCards(db *pgxpool.Pool) (err error) {
	query := "TRUNCATE TABLE cards"

	_, err = db.Exec(context.Background(), query)

	return
}
