package db

import (
	"context"
	"fmt"
	"os" // Added to read environment variables
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	// 1. Try to get the connection string from the environment (for Docker)
	connStr := os.Getenv("DATABASE_URL")

	// 2. If it's empty (running locally on Windows), use the default fallback
	if connStr == "" {
		connStr = "postgres://user:password@127.0.0.1:5432/sentrydb?sslmode=disable"
	}

	fmt.Printf("🔌 Connecting to: %s\n", connStr)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return pool, nil
}