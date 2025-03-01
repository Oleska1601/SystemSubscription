package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (pgRepo *PostgresRepo) GetSubcriptionType(ctx context.Context, subscriptionTypeName string) (entity.SubscriptionType, error) {
	sql, args, err := pgRepo.db.Builder.Select("subscription_type_id", "type_name", "duration", "price").
		From("subscription_types").Where(squirrel.Eq{"type_name": subscriptionTypeName}).ToSql()
	if err != nil {
		return entity.SubscriptionType{}, fmt.Errorf("repo-repoAPI GetSubcriptionTypeID pgRepo.db.Builder.Select: %w", err)
	}
	var subscriptionType entity.SubscriptionType
	err = pgRepo.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&subscriptionType.ID,
		&subscriptionType.TypeName,
		&subscriptionType.Duration,
		&subscriptionType.Price,
	)
	if err != nil {
		return entity.SubscriptionType{}, fmt.Errorf("repo-repoAPI GetSubcriptionTypeID pgRepo.db.Pool.QueryRow.Scan: %w", err)
	}
	return subscriptionType, nil
}

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
