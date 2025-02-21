package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (pgRepo *PostgresRepo) GetSubscriptionTypes(ctx context.Context) ([]entity.SubscriptionType, error) {
	subscriptionTypes := []entity.SubscriptionType{}
	sql, args, err := pgRepo.db.Builder.Select("subscription_type_id", "type_name", "duration", "price").From("subscription_types").ToSql()
	if err != nil {
		return nil, fmt.Errorf("repo-repoAPI GetSubscriptionTypes, pgRepo.db.Builder.Select: %w", err)
	}
	rows, err := pgRepo.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repo-repoAPI GetSubscriptionTypes, pgRepo.db.Pool.Query: %w", err)
	}
	for rows.Next() {
		var subscriptionType entity.SubscriptionType
		err = rows.Scan(&subscriptionType.ID, &subscriptionType.TypeName, &subscriptionType.Duration, &subscriptionType.Price)
		if err != nil {
			return nil, fmt.Errorf("repo-repoAPI GetSubscriptionTypes, rows.Scan: %w", err)
		}
		subscriptionTypes = append(subscriptionTypes, subscriptionType)
	}
	return subscriptionTypes, nil
}

func (pgRepo *PostgresRepo) GetLastSubscription(ctx context.Context, userID int64) (entity.Subscription, error) {
	sql, args, err := pgRepo.db.Builder.Select("subscription_id", "status", "start_date", "end_date", "subscription_type_id").From("subscriptions").Where(squirrel.Eq{"user_id": userID}).OrderBy("end_date desc").Limit(1).ToSql()
	if err != nil {
		return entity.Subscription{}, fmt.Errorf("repo-repoAPI GetLastSubscription pgRepo.db.Builder.Select: %w", err)
	}
	row := pgRepo.db.Pool.QueryRow(ctx, sql, args...)
	var subscription entity.Subscription
	err = row.Scan(&subscription.ID, &subscription.Status, &subscription.StartDate, &subscription.EndDate, &subscription.SubscriptionTypeID)
	if err != nil {
		return entity.Subscription{}, fmt.Errorf("repo-repoAPI GetLastSubscription row.Scan: %w", err)
	}
	return subscription, nil
}
