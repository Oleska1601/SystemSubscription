package repo

import (
	"SystemSubscription/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

func (pgRepo *PostgresRepo) GetPayment(ctx context.Context, paymentToken string) (entity.Payment, error) {
	sql, args, err := pgRepo.db.Builder.Select("payment_id", "token", "subscription_type_name", "amount",
		"start_time", "end_time", "payment_status_id", "user_id").
		From("payments").Where(squirrel.Eq{"token": paymentToken}).ToSql()
	if err != nil {
		return entity.Payment{}, fmt.Errorf("repo GetPayment pgRepo.db.Builder.Select: %w", err)
	}
	var payment entity.Payment
	err = pgRepo.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&payment.ID,
		&payment.Token,
		&payment.SubscriptionTypeName,
		&payment.Amount,
		&payment.StartTime,
		&payment.EntTime,
		&payment.PaymentStatusID,
		&payment.UserID)
	if err != nil {
		return entity.Payment{}, fmt.Errorf("repo GetPayment pgRepo.db.Pool.QueryRow.Scan: %w", err)
	}
	return payment, nil
}

func (pgRepo *PostgresRepo) InsertPayment(ctx context.Context, payment *entity.Payment) error {
	sql, args, err := pgRepo.db.Builder.Insert("payments").
		Columns("token", "subscription_type_name", "amount", "start_time", "end_time", "payment_status_id", "user_id").
		Values(payment.Token, payment.SubscriptionTypeName, payment.Amount,
			payment.StartTime, payment.EntTime, payment.PaymentStatusID, payment.UserID).ToSql()
	if err != nil {
		return fmt.Errorf("repo InsertPayment pgRepo.db.Builder.Insert: %w", err)
	}
	_, err = pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo InsertPayment pgRepo.db.Pool.Exec: %w", err)
	}
	return nil
}

func (pgRepo *PostgresRepo) UpdatePayment(ctx context.Context, paymentID int64, newPaymentStatusID int) error {
	sql, args, err := pgRepo.db.Builder.Update("payments").
		Set("payment_status_id", newPaymentStatusID).
		Where(squirrel.Eq{"payment_id": paymentID}).ToSql()
	if err != nil {
		return fmt.Errorf("repo UpdatePayment pgRepo.db.Builder.Update: %w", err)
	}
	_, err = pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo UpdatePayment pgRepo.db.Pool.Exec: %w", err)
	}
	return nil
}

func (pgRepo *PostgresRepo) UpdatePayments(ctx context.Context, newPaymentStatusID int, timeNow time.Time) error {
	sql, args, err := pgRepo.db.Builder.Update("payments").
		Set("payment_status_id", newPaymentStatusID).
		Where(squirrel.Gt{"end_time": timeNow}).ToSql()
	if err != nil {
		return fmt.Errorf("repo UpdatePayments pgRepo.db.Builder.Update: %w", err)
	}
	_, err = pgRepo.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repo UpdatePayments pgRepo.db.Pool.Exec: %w", err)
	}
	return nil
}
