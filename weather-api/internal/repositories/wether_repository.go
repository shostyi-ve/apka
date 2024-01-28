package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type WeatherRepository struct {
	db *pgx.Conn
}

func New(conn *pgx.Conn) *WeatherRepository {
	return &WeatherRepository{db: conn}
}

func (r *WeatherRepository) Close(ctx context.Context) error {
	if err := r.db.Close(ctx); err != nil {
		return fmt.Errorf("storage close: %w", err)
	}

	return nil
}
