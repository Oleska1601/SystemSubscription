package controller

import "SystemSubscription/internal/entity"

type UseCaseTestWithNoErr struct{}

func NewUseCaseTestWithNoErr() *UseCaseTestWithNoErr {
	return &UseCaseTestWithNoErr{}
}

func (u *UseCaseTestWithNoErr) ActivateSubscription(paymentToken string) error {
	return nil
}

func (u *UseCaseTestWithNoErr) AddPayment(subscriptionType *entity.SubscriptionType, userID int64) (string, error) {
	return "", nil
}
func (u *UseCaseTestWithNoErr) AddSubscription(subscriptionTypeName string, userID int64) error {
	return nil
}
func (u *UseCaseTestWithNoErr) GenerateHash(password string, salt string) string {
	return ""
}
func (u *UseCaseTestWithNoErr) GeneratePaymentToken() string {
	return ""
}
func (u *UseCaseTestWithNoErr) GenerateSalt() string {
	return ""
}
func (u *UseCaseTestWithNoErr) GetLastSubscription(userID int64) (entity.Subscription, error) {
	return entity.Subscription{}, nil
}
func (u *UseCaseTestWithNoErr) GetLastSubscriptionInfo(userID int64) (entity.SubscriptionInfo, error) {
	return entity.SubscriptionInfo{}, nil
}
func (u *UseCaseTestWithNoErr) GetNews(userID int64) (entity.News, error) {
	return entity.News{}, nil
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
func (u *UseCaseTestWithNoErr) IsPasswordCorrect(inputPassword string, loginUser entity.User) bool {
	return true
}
func (u *UseCaseTestWithNoErr) IsPaymentActive(payment entity.Payment) bool {
	return true
}
func (u *UseCaseTestWithNoErr) IsSubscriptionStatusActive(subscription entity.Subscription) bool {
	return true
}
func (u *UseCaseTestWithNoErr) IsUserExists(loginUser entity.User) bool {
	return true
}
func (u *UseCaseTestWithNoErr) LoginUser(user *entity.User) (entity.User, error) {
	return entity.User{}, nil
}
func (u *UseCaseTestWithNoErr) RegisterUser(user *entity.User) error {
	return nil
}
func (u *UseCaseTestWithNoErr) SetNewSubscription(userID int64, subscriptionType entity.SubscriptionType) (string, error) {
	return "", nil
}
func (u *UseCaseTestWithNoErr) UpdatePayment(paymentID int64, newPaymentStatusID int) error {
	return nil
}
