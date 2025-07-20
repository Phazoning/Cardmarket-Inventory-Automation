package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func FetchDB() (db *pgxpool.Pool, err error) {
	err = godotenv.Load("../.env")

	if err != nil {
		return
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
	db, err = pgxpool.New(context.Background(), connStr)
	return
}
