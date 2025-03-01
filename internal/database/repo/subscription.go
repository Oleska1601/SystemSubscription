package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

func (pgRepo *PostgresRepo) GetLastSubscription(ctx context.Context, userID int64) (entity.Subscription, error) {
	sql, args, err := pgRepo.db.Builder.Select("subscription_id", "status", "start_date", "end_date", "subscription_type_id").
		From("subscriptions").
		Where(squirrel.Eq{"user_id": userID}).OrderBy("end_date desc").Limit(1).ToSql()
	if err != nil {
		return entity.Subscription{}, fmt.Errorf("repo-repoAPI GetLastSubscription pgRepo.db.Builder.Select: %w", err)
	}
	row := pgRepo.db.Pool.QueryRow(ctx, sql, args...)
	var subscription entity.Subscription
	err = row.Scan(&subscription.ID, &subscription.Status, &subscription.StartDate, &subscription.EndDate, &subscription.SubscriptionTypeID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Subscription{}, nil
		}
		return entity.Subscription{}, fmt.Errorf("repo-repoAPI GetLastSubscription row.Scan: %w", err)
	}
	return subscription, nil
}

func (pgRepo *PostgresRepo) InsertSubscription(ctx context.Context, subcription *entity.Subscription) (int64, error) {
	sql, args, err := pgRepo.db.Builder.Insert("subscriptions").
		Columns("status", "start_date", "end_date", "subscription_type_id", "user_id").
		Values(subcription.Status, subcription.StartDate, subcription.EndDate, subcription.SubscriptionTypeID, subcription.UserID).
		Suffix("RETURNING subscription_id").ToSql()
	if err != nil {
		return 0, fmt.Errorf("repo-repoAPI InsertSubscription pgRepo.db.Builder.Insert: %w", err)
	}
	var subscriptionID int64
	err = pgRepo.db.Pool.QueryRow(ctx, sql, args...).Scan(&subscriptionID)
	if err != nil {
		return 0, fmt.Errorf("repo-repoAPI InsertSubscription pgRepo.db.Pool.QueryRow.Scan: %w", err)
	}
	return subscriptionID, nil
}
