package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

// получить пользователя по текущ логину и возвратить его данные иначе пустую структуру
func (pgRepo *PostgresRepo) GetUser(ctx context.Context, userLogin string) (entity.User, error) {
	sql, args, err := pgRepo.db.Builder.Select("user_id", "login", "password_hash", "password_salt").From("users").Where(squirrel.Eq{"login": userLogin}).ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("repo-user GetUser pgRepo.db.Builder.Select: %w", err)
	}
	var user entity.User
	err = pgRepo.db.Pool.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Login, &user.PasswordHash, &user.PasswordSalt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, nil
		}
		return entity.User{}, fmt.Errorf("repo-user GetUser row.Scan: %w", err)
	}
	return user, nil
}

func (pgRepo *PostgresRepo) InsertUser(ctx context.Context, user *entity.User) error {
	sql, args, err := pgRepo.db.Builder.Insert("users").Columns("login", "password_hash", "password_salt").Values(user.Login, user.PasswordHash, user.PasswordSalt).ToSql()
	if err != nil {
		return fmt.Errorf("repo-user InsertUser pgRepo.db.Builder.Insert: %w", err)
	}
	_, err = pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("database-user InsertUser pgRepo.db.Pool.Exec: %w", err)
	}
	return nil
}
