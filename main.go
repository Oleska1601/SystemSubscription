package main

import (
	"SystemSubscription/config"
	"SystemSubscription/internal/controller"
	"SystemSubscription/internal/database/repo"
	"SystemSubscription/internal/usecase"
	"SystemSubscription/pkg/logger"
	"SystemSubscription/pkg/postgres"
	"log/slog"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("main - config.New", slog.Any("error", err))
		return
	}
	l := logger.New(cfg.Logger.Level)
	pg, err := postgres.New(l, cfg.Postgres.PgUrl, postgres.MaxPoolSize(cfg.PoolMax))
	if err != nil {
		l.Error("main - postgres.New", slog.Any("error", err))
		return
	}
	pgRepo := repo.New(pg)
	if err = pgRepo.CreateTables(); err != nil {
		l.Error("main - db.CreateTables", slog.Any("error", err))
		return
	}
	l.Info("connected to db")
	u := usecase.New(pgRepo, l)
	s := controller.New(u, l)
	s.Run(cfg.HTTP.Port)
}
