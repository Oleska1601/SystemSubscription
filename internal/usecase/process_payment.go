package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

func (u *Usecase) GeneratePaymentToken() string {
	paymentTokenBytes := make([]byte, 64)
	rand.Read(paymentTokenBytes)
	return base64.StdEncoding.EncodeToString(paymentTokenBytes)
}

func (u *Usecase) IsPaymentActive(payment entity.Payment) bool {
	return payment.EntTime.After(time.Now())
}

func (u *Usecase) GetPayment(paymentToken string) (entity.Payment, error) {
	ctx, _ := context.WithCancel(context.Background())
	payment, err := u.pgRepo.GetPayment(ctx, paymentToken)
	if err != nil {
		return entity.Payment{}, fmt.Errorf("u.pgRepo.GetPayment %w", err)
	}
	return payment, nil
}

func (u *Usecase) AddPayment(subscriptionType *entity.SubscriptionType, userID int64) (string, error) {
	ctx, _ := context.WithCancel(context.Background())
	startTime := time.Now()
	endTime := startTime.Add(time.Second * 100)
	//startTimeFormat := startTime.Format("2006-01-02T15:04:05Z")
	//endTimeFormat = endTime.Format("2006-01-02T15:04:05Z")
	paymentToken := u.GeneratePaymentToken()
	payment := entity.Payment{
		Token:                paymentToken,
		SubscriptionTypeName: subscriptionType.TypeName,
		Amount:               subscriptionType.Price,
		StartTime:            startTime,
		EntTime:              endTime,
		// p.s по хорошему здесь делаем запрос в бд и получаем payment_status_id где payment_status=in process
		PaymentStatusID: 2,
		UserID:          userID,
	}
	err := u.pgRepo.InsertPayment(ctx, &payment)
	if err != nil {
		return "", fmt.Errorf("u.pgRepo.InsertPayment %w", err)
	}
	return paymentToken, nil
}

func (u *Usecase) UpdatePayment(paymentID int64, newPaymentStatusID int) error {
	ctx, _ := context.WithCancel(context.Background())
	err := u.pgRepo.UpdatePayment(ctx, paymentID, newPaymentStatusID)
	if err != nil {
		return fmt.Errorf("u.pgRepo.UpdatePayment %w", err)
	}
	return nil
}
