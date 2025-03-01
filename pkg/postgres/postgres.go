package postgres

import (
	"SystemSubscription/pkg/logger"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	_defaultMaxPoolSize     = 1
	_defaultMaxConnAttempts = 10
	_defaultMaxConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	Builder      squirrel.StatementBuilderType
	Pool         *pgxpool.Pool
}

func New(logger *logger.Logger, url string, options ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultMaxConnAttempts,
		connTimeout:  _defaultMaxConnTimeout,
	}
	for _, opt := range options {
		opt(pg)
	}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	//создает конфигурацию для пула соединений - парсит url и создает pgxpool.COnfig
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ParseConfig: %w", err)
	}
	poolConfig.MaxConns = int32(pg.maxPoolSize)
	for pg.connAttempts > 0 {
		logger.Info("Postgres is trying to connect", slog.Int("conntection attempts left:", pg.connAttempts))
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		pg.connAttempts--
	}
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.ConnectConfig: %w", err)
	}
	return pg, nil
}

func (pg *Postgres) Close() {
	if pg.Pool != nil {
		pg.Pool.Close()
	}
}
