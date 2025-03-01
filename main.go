package main

import (
	"SystemSubscription/config"
	_ "SystemSubscription/docs"
	"SystemSubscription/internal/controller"
	"SystemSubscription/internal/database/repo"
	"SystemSubscription/internal/usecase"
	"SystemSubscription/pkg/logger"
	"SystemSubscription/pkg/postgres"
	"log/slog"
)

// @title System Of Subscriptions API
// @version 1.0
// @description API for Golang Project System of Subscriptions
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /
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
