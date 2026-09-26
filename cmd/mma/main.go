package main

import (
	"context"
	"log/slog"
	postgres "mma/db/postgres"
	config "mma/internal/config"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		logger.Error("ошибка при загрузке конфига", "error", err)
		return
	}
	ctx := context.Background()
	db, err := postgres.NewPostgres(ctx, cfg.PostgresURL)
	if err != nil {
		logger.Error("ошибка при подключении БД", "error", err)
		return
	}
	defer db.Close()
	migrator, err := postgres.NewMigrator("", cfg.PostgresURL)
	if err != nil {
		logger.Error("ошибка при подключении мигратора", "error", err)
		return
	}
	if err := migrator.Up(); err != nil {
		logger.Error("ошибка при поднятии мигратора", "error", err)
		return
	}
	defer migrator.Close()
}