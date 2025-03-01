package postgres

import "time"

type Option func(*Postgres)

// возвращают анонимную функцию, которая принимает указатель на `Postgres` и изменяет соответствующие поля
func MaxPoolSize(size int) Option {
	return func(p *Postgres) {
		p.maxPoolSize = size
	}
}

func MaxConnAttempts(attempts int) Option {
	return func(p *Postgres) {
		p.connAttempts = attempts
	}
}

func MaxConnTimeout(timeout time.Duration) Option {
	return func(p *Postgres) {
		p.connTimeout = timeout
	}
}
