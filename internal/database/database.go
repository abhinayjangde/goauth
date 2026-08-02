package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(databaseURL string) (*pgxpool.Pool, error) {

	poop, err := pgxpool.New(context.Background(), databaseURL)

	if err != nil {
		return nil, err
	}

	if err := poop.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Database Connected Successfully")

	return poop, nil
}

func CheckConnection(pool *pgxpool.Pool) (string, error) {
	var version string

	err := pool.QueryRow(
		context.Background(),
		"SELECT version()",
	).Scan(&version)

	if err != nil {
		return "", err
	}

	return version, nil
}
