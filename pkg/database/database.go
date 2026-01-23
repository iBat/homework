package database

import (
	"context"
	"iBat/homework/config"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDbPool(config *config.DatabaseConfig, logger *slog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)

	if err != nil {
		logger.Error("Unable to connect to DB")
		panic(err)
	}
	logger.Info("Connected to DB")
	return dbpool
}
