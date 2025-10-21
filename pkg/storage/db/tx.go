package db

import (
	"context"

	"echo.go.dev/pkg/storage/db/dbx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// WithTx Helper to run a handler function in a transaction
func WithTx(ctx context.Context, pool *pgxpool.Pool, fn func(pgx.Tx, *dbx.Queries) error) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := dbx.New(pool).WithTx(tx)
	if err := fn(tx, q); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
