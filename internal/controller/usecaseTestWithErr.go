package controller

import (
	"SystemSubscription/internal/entity"
	"errors"
)

type UseCaseTestWithErr struct {
	err error
}

func NewUseCaseTestWithErr() *UseCaseTestWithErr {
	return &UseCaseTestWithErr{
		err: errors.New("test"),
	}
}

func (u *UseCaseTestWithErr) AddPayment(subscriptionType *entity.SubscriptionType, userID int64) (string, error) {
	return "", u.err
}

func (u *UseCaseTestWithErr) AddSubscription(subscriptionTypeName string, userID int64) (int64, error) {
	return 0, u.err
}

func (u *UseCaseTestWithErr) GetLastSubscription(userID int64) (entity.Subscription, error) {
	return entity.Subscription{}, u.err
}
func (u *UseCaseTestWithErr) GetPayment(paymentToken string) (entity.Payment, error) {
	return entity.Payment{}, u.err
}

func (u *UseCaseTestWithErr) GetSubscriptionTypes() ([]entity.SubscriptionType, error) {
	return []entity.SubscriptionType{}, u.err
}
func (u *UseCaseTestWithErr) IsLastSubscription(subscription entity.Subscription) bool {
	return false
}

func (u *UseCaseTestWithErr) IsPaymentActive(payment entity.Payment) bool {
	return false
}

func (u *UseCaseTestWithErr) IsSubscriptionStatusActive(subscription entity.Subscription) bool {
	return false
}

func (u *UseCaseTestWithErr) LoginUser(user *entity.User) (entity.User, error) {
	return entity.User{}, u.err
}

func (u *UseCaseTestWithErr) RegisterUser(user *entity.User) error {
	return u.err
}

func (u *UseCaseTestWithErr) UpdatePayment(paymentID int64, newPaymentStatus string) error {
	return u.err
}
