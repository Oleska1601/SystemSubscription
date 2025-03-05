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

func (u *UseCaseTestWithErr) ActivateSubscription(paymentToken string) error {
	return u.err
}

func (u *UseCaseTestWithErr) AddPayment(subscriptionType *entity.SubscriptionType, userID int64) (string, error) {
	return "", u.err
}
func (u *UseCaseTestWithErr) AddSubscription(subscriptionTypeName string, userID int64) error {
	return u.err
}
func (u *UseCaseTestWithErr) GenerateHash(password string, salt string) string {
	return ""
}
func (u *UseCaseTestWithErr) GeneratePaymentToken() string {
	return ""
}
func (u *UseCaseTestWithErr) GenerateSalt() string {
	return ""
}
func (u *UseCaseTestWithErr) GetLastSubscription(userID int64) (entity.Subscription, error) {
	return entity.Subscription{}, u.err
}
func (u *UseCaseTestWithErr) GetLastSubscriptionInfo(userID int64) (entity.SubscriptionInfo, error) {
	return entity.SubscriptionInfo{}, u.err
}
func (u *UseCaseTestWithErr) GetNews(userID int64) (entity.News, error) {
	return entity.News{}, u.err
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
func (u *UseCaseTestWithErr) IsPasswordCorrect(inputPassword string, loginUser entity.User) bool {
	return false
}
func (u *UseCaseTestWithErr) IsPaymentActive(payment entity.Payment) bool {
	return false
}
func (u *UseCaseTestWithErr) IsSubscriptionStatusActive(subscription entity.Subscription) bool {
	return false
}
func (u *UseCaseTestWithErr) IsUserExists(loginUser entity.User) bool {
	return false
}
func (u *UseCaseTestWithErr) LoginUser(user *entity.User) (entity.User, error) {
	return entity.User{}, u.err
}
func (u *UseCaseTestWithErr) RegisterUser(user *entity.User) error {
	return u.err
}
func (u *UseCaseTestWithErr) SetNewSubscription(userID int64, subscriptionType entity.SubscriptionType) (string, error) {
	return "", u.err
}
func (u *UseCaseTestWithErr) UpdatePayment(paymentID int64, newPaymentStatusID int) error {
	return u.err
}
