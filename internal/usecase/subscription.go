package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) SetNewSubscription(userID int64, subscriptionType entity.SubscriptionType) (string, error) {
	lastSubscription, err := u.GetLastSubscription(userID)
	if err != nil {
		u.logger.Error("usecase SetNewSubscription u.GetLastSubscription", slog.Any("error", err))
		return "", errors.New("error of setting new subscription")
	}
	if u.IsSubscriptionStatusActive(lastSubscription) {
		u.logger.Error("usecase SetNewSubscription u.IsSubscriptionStatusActive",
			slog.String("msg", "current subscription is active"))
		return "", errors.New("error of setting new subscription")
	}
	paymentToken, err := u.AddPayment(&subscriptionType, userID)
	if err != nil {
		u.logger.Error("usecase SetNewSubscription u.AddPayment", slog.Any("error", err))
		return "", errors.New("error of setting new subscription")
	}
	return paymentToken, nil
}

func (u *Usecase) ActivateSubscription(paymentToken string) error {
	payment, err := u.GetPayment(paymentToken)
	if err != nil {
		u.logger.Error("usecase ActivateSubscription u.GetPayment", slog.Any("error", err))
		return errors.New("error of activating subcription")
	}
	if !u.IsPaymentActive(payment) {
		//какую ошибку возращать: payment token has been expired и устанавливать статус
		// forbidden или как и везде error of activating subcription internal server error
		u.logger.Error("usecase ActivateSubscription u.IsPaymentActive",
			slog.String("msg", "payment token has been expired"))
		return errors.New("error of activating subcription")
	}
	// p.s по хорошему здесь делаем запрос в бд и получаем payment_status_id где payment_status=success
	newPaymentStatusID := 1 // 1-success
	err = u.UpdatePayment(payment.ID, newPaymentStatusID)
	if err != nil {
		u.logger.Error("usecase ActivateSubscription u.UpdatePayment", slog.Any("error", err))
		return errors.New("error of activating subcription")
	}
	err = u.AddSubscription(payment.SubscriptionTypeName, payment.UserID)
	if err != nil {
		u.logger.Error("usecase ActivateSubscription u.AddSubscription", slog.Any("error", err))
		return errors.New("error of activating subcription")
	}
	return nil
}

func (u *Usecase) GetLastSubscriptionInfo(userID int64) (entity.SubscriptionInfo, error) {
	ctx, _ := context.WithCancel(context.Background())
	lastSubscriptionInfo, err := u.pgRepo.GetLastSubscriptionInfo(ctx, userID)
	if err != nil {
		u.logger.Error("usecase GetLastSubscriptionInfo u.pgRepo.GetLastSubscriptionInfo", slog.Any("error", err))
		return entity.SubscriptionInfo{}, errors.New("error of getting last subscription info")
	}
	return lastSubscriptionInfo, nil
}
