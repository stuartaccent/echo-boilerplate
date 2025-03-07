package db

import (
	"context"
	"echo.go.dev/pkg/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConn(ctx context.Context) (*pgx.Conn, error) {
	cfg := config.GetConfig()

	conn, err := pgx.Connect(ctx, cfg.Database.URL().String())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetPool(ctx context.Context) (*pgxpool.Pool, error) {
	cfg := config.GetConfig()

	pool, err := pgxpool.New(ctx, cfg.Database.URL().String())
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
