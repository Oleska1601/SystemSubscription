package usecase

import (
	"SystemSubscription/internal/database/repo"
	"SystemSubscription/pkg/logger"
)

type Usecase struct {
	pgRepo *repo.PostgresRepo
	logger *logger.Logger
}

func New(pgRepo *repo.PostgresRepo, l *logger.Logger) *Usecase {
	return &Usecase{
		pgRepo: pgRepo,
		logger: l,
	}
}
