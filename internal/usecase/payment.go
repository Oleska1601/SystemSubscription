package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"log/slog"
	"time"
)

func (u *Usecase) GetPayment(paymentToken string) (entity.Payment, error) {
	ctx, _ := context.WithCancel(context.Background())
	payment, err := u.pgRepo.GetPayment(ctx, paymentToken)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI ActivateSubscription u.pgRepo.GetPayment", slog.Any("error", err))
		return entity.Payment{}, errors.New("error of getting payment")
	}
	return payment, nil
}

func (u *Usecase) AddPayment(subscriptionType *entity.SubscriptionType, userID int64) (string, error) {
	ctx, _ := context.WithCancel(context.Background())
	startTime := time.Now()
	endTime := startTime.Add(time.Second * 100)
	//startTimeFormat := startTime.Format("2006-01-02T15:04:05Z")
	//endTimeFormat = endTime.Format("2006-01-02T15:04:05Z")
	paymentToken := GeneratePaymentToken()
	payment := entity.Payment{
		Token:                paymentToken,
		SubscriptionTypeName: subscriptionType.TypeName,
		Amount:               subscriptionType.Price,
		StartTime:            startTime,
		EntTime:              endTime,
		Status:               "token is generated",
		UserID:               userID,
	}
	err := u.pgRepo.InsertPayment(ctx, &payment)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI AddPayment u.pgRepo.InsertPayment", slog.Any("error", err))
		return "", errors.New("error of adding payment")
	}
	return paymentToken, nil
}

func (u *Usecase) UpdatePayment(paymentID int64, newPaymentStatus string) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.pgRepo.UpdatePayment(ctx, paymentID, newPaymentStatus)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI UpdatePayment u.pgRepo.UpdatePayment", slog.Any("error", err))
		return errors.New("error of updating payment")
	}
	return nil
}

func (u *Usecase) IsPaymentActive(payment entity.Payment) bool {
	return payment.EntTime.After(time.Now())
}
