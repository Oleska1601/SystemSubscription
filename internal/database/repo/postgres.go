package repo

import (
	"SystemSubscription/pkg/postgres"
	"context"
)

type PostgresRepo struct {
	db *postgres.Postgres
}

func New(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{db: pg}
}

func (pgRepo *PostgresRepo) CreateTables() error {
	queries := []string{
		`create table users (
		user_id serial primary key,
		login varchar(100) not null unique,
		password_hash text not null,
		password_salt text not null,
		secret text not null
	);`,
		`create table subscription_types (
		subscription_type_id serial primary key,
		type_name varchar(50) not null unique,
		duration integer not null unique,
		price integer not null
	);`,
		`create table subscriptions (
		subscription_id serial primary key,
		status varchar(20) not null,
		start_date timestamp with time zone not null,
		end_date timestamp with time zone not null,
		subscription_type_id integer references subscription_types(subscription_type_id) not null,
		user_id integer references users(user_id) not null
	);`,
		`create table payments (
		payment_id serial primary key,
		status varchar(30) not null,
		start_time timestamp with time zone not null,
		end_time timestamp with time zone not null,
		subscription_id integer references subscriptions(subscription_id) not null,
		user_id integer references users(user_id) not null
	);`}
	for _, query := range queries {
		_, err := pgRepo.db.Pool.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}
