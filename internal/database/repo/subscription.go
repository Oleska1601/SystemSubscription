package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

func (pgRepo *PostgresRepo) GetLastSubscription(ctx context.Context, userID int64) (entity.Subscription, error) {
	sql, args, err := pgRepo.db.Builder.Select("subscription_id",
		"start_date", "end_date", "subscription_type_id", "subscription_status_id").
		From("subscriptions").
		Where(squirrel.Eq{"user_id": userID}).OrderBy("end_date desc").Limit(1).ToSql()
	if err != nil {
		return entity.Subscription{}, fmt.Errorf("repo GetLastSubscription pgRepo.db.Builder.Select: %w", err)
	}
	row := pgRepo.db.Pool.QueryRow(ctx, sql, args...)
	var subscription entity.Subscription
	err = row.Scan(&subscription.ID, &subscription.StartDate,
		&subscription.EndDate, &subscription.SubscriptionTypeID, &subscription.SubscriptionStatusID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Subscription{}, nil
		}
		return entity.Subscription{}, fmt.Errorf("repo GetLastSubscription row.Scan: %w", err)
	}
	return subscription, nil
}

func (pgRepo *PostgresRepo) GetLastSubscriptionInfo(ctx context.Context, userID int64) (entity.SubscriptionInfo, error) {
	sql, args, err := pgRepo.db.Builder.Select("s.subscription_id",
		"s.start_date", "s.end_date", "st.type_name", "ss.subscription_status").
		From("subscriptions as s").
		Join("subscription_status as ss on ss.subscription_status_id=s.subscription_status_id").
		Join("subscription_types as st on st.subscription_type_id=s.subscription_type_id").
		Where(squirrel.Eq{"user_id": userID}).OrderBy("end_date desc").Limit(1).ToSql()
	if err != nil {
		return entity.SubscriptionInfo{}, fmt.Errorf("repo GetLastSubscriptionInfo pgRepo.db.Builder.Select: %w", err)
	}
	row := pgRepo.db.Pool.QueryRow(ctx, sql, args...)
	var subscriptionInfo entity.SubscriptionInfo
	err = row.Scan(&subscriptionInfo.ID, &subscriptionInfo.StartDate,
		&subscriptionInfo.EndDate, &subscriptionInfo.SubscriptionTypeName, &subscriptionInfo.SubscriptionStatus)
	if err != nil {
		return entity.SubscriptionInfo{}, fmt.Errorf("repo GetLastSubscriptionInfo row.Scan: %w", err)
	}
	return subscriptionInfo, nil
}

func (pgRepo *PostgresRepo) InsertSubscription(ctx context.Context, subcription *entity.Subscription) error {
	sql, args, err := pgRepo.db.Builder.Insert("subscriptions").
		Columns("start_date", "end_date", "subscription_type_id", "subscription_status_id", "user_id").
		Values(subcription.StartDate, subcription.EndDate, subcription.SubscriptionTypeID,
			subcription.SubscriptionStatusID, subcription.UserID).ToSql()
	if err != nil {
		return fmt.Errorf("repo InsertSubscription pgRepo.db.Builder.Insert: %w", err)
	}
	pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo InsertSubscription pgRepo.db.Pool.QueryRow.Scan: %w", err)
	}
	return nil
}

func (pgRepo *PostgresRepo) UpdateSubscriptions(ctx context.Context, newSubscriptionStatusID int, timeNow time.Time) error {
	sql, args, err := pgRepo.db.Builder.Update("subscriptions").
		Set("subscription_status_id", newSubscriptionStatusID).
		Where(squirrel.Gt{"end_date": timeNow}).ToSql()
	if err != nil {
		return fmt.Errorf("repo UpdateSubscriptions pgRepo.db.Builder.Update: %w", err)
	}
	_, err = pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo UpdateSubscriptions pgRepo.db.Pool.Exec: %w", err)
	}
	return nil
}
