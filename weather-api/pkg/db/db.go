package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func GetConnect(ctx context.Context, config *Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, config.GetConnectString())
	if err != nil {
		return nil, fmt.Errorf("storage connect: %w", err)
	}

	return conn, nil
}
