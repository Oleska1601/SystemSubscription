package repo

import (
	"SystemSubscription/pkg/postgres"
	"context"
	"os"
	"strings"
)

type PostgresRepo struct {
	db *postgres.Postgres
}

func New(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{db: pg}
}

func (pgRepo *PostgresRepo) CreateTables() error {
	queriesBytes, err := os.ReadFile("./internal/database/repo/sql/query.sql")
	if err != nil {
		return err
	}
	queries := strings.Split(string(queriesBytes), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query) //убрать лишние пробелы
		if query == "" {
			continue
		}
		_, err := pgRepo.db.Pool.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}
