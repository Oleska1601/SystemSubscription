package controller

import "SystemSubscription/internal/entity"

type UseCaseTestWithNoErr struct{}

func NewUseCaseTestWithNoErr() *UseCaseTestWithNoErr {
	return &UseCaseTestWithNoErr{}
}

func (u *UseCaseTestWithNoErr) AddPayment(subscriptionType *entity.SubscriptionType, userID int64) (string, error) {
	return "", nil
}

func (u *UseCaseTestWithNoErr) AddSubscription(subscriptionTypeName string, userID int64) (int64, error) {
	return 0, nil
}

func (u *UseCaseTestWithNoErr) GetLastSubscription(userID int64) (entity.Subscription, error) {
	return entity.Subscription{}, nil
}
func (u *UseCaseTestWithNoErr) GetPayment(paymentToken string) (entity.Payment, error) {
	return entity.Payment{}, nil
}

func (u *UseCaseTestWithNoErr) GetSubscriptionTypes() ([]entity.SubscriptionType, error) {
	return []entity.SubscriptionType{}, nil
}
func (u *UseCaseTestWithNoErr) IsLastSubscription(subscription entity.Subscription) bool {
	return true
}

func (u *UseCaseTestWithNoErr) IsPaymentActive(payment entity.Payment) bool {
	return true
}

func (u *UseCaseTestWithNoErr) IsSubscriptionStatusActive(subscription entity.Subscription) bool {
	return true
}

func (u *UseCaseTestWithNoErr) LoginUser(user *entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (u *UseCaseTestWithNoErr) RegisterUser(user *entity.User) error {
	return nil
}

func (u *UseCaseTestWithNoErr) UpdatePayment(paymentID int64, newPaymentStatus string) error {
	return nil
}
