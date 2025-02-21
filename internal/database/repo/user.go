package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (pgRepo *PostgresRepo) InsertUser(ctx context.Context, user *entity.User) error {
	sql, args, err := pgRepo.db.Builder.Insert("users").Columns("login", "password_hash", "password_salt", "secret").Values(user.Login, user.PasswordHash, user.PasswordSalt, user.Secret).ToSql()
	if err != nil {
		return fmt.Errorf("repo-user InsertUser pgRepo.db.Builder.Insert: %w", err)
	}
	_, err = pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("database-user InsertUser pgRepo.db.Pool.Exec: %w", err)
	}
	return nil
}

// получить пользователя по текущ логину и возвратить его данные иначе пустую структуру
func (pgRepo *PostgresRepo) GetUser(ctx context.Context, userLogin string) (entity.User, error) {
	sql, args, err := pgRepo.db.Builder.Select("user_id").From("users").Where(squirrel.Eq{"login": userLogin}).ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("repo-user GetUser pgRepo.db.Builder.Select: %w", err)
	}
	row := pgRepo.db.Pool.QueryRow(ctx, sql, args...)
	var user entity.User
	err = row.Scan(&user.ID, &user.Login, &user.PasswordHash, &user.PasswordSalt, &user.Secret)
	if err != nil {
		return entity.User{}, fmt.Errorf("repo-user GetUser row.Scan: %w", err)
	}
	return user, nil
}
