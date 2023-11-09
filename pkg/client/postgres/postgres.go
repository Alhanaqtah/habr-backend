package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func NewClient(user, password, host, port, db string) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		user,
		password,
		host,
		port,
		db)

	ctx := context.WithoutCancel(context.Background())

	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("failed to create conn pool: %v", err.Error())
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("failed connect to repository: %v", err.Error())
	}

	return pool, nil
}
