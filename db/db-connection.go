package db

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DbConnection(ctx context.Context) *pgxpool.Pool {
	dbURL := os.Getenv("POSTGRESQL_URL")
	if dbURL == "" {
		log.Fatal("POSTGRESQL_URL config missing")
	}
	dbURL = strings.Trim(dbURL, `"'`)
	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}

	return pool
}
