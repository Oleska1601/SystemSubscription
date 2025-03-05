package controller

import "SystemSubscription/internal/entity"

type (
	UseCaseInterface interface {
		ActivateSubscription(string) error
		AddPayment(*entity.SubscriptionType, int64) (string, error)
		AddSubscription(string, int64) error
		GenerateHash(string, string) string
		GeneratePaymentToken() string
		GenerateSalt() string
		GetLastSubscription(int64) (entity.Subscription, error)
		GetLastSubscriptionInfo(int64) (entity.SubscriptionInfo, error)
		GetNews(int64) (entity.News, error)
		GetPayment(string) (entity.Payment, error)
		GetSubscriptionTypes() ([]entity.SubscriptionType, error)
		IsLastSubscription(entity.Subscription) bool
		IsPasswordCorrect(string, entity.User) bool
		IsPaymentActive(entity.Payment) bool
		IsSubscriptionStatusActive(entity.Subscription) bool
		IsUserExists(entity.User) bool
		LoginUser(*entity.User) (entity.User, error)
		RegisterUser(*entity.User) error
		SetNewSubscription(int64, entity.SubscriptionType) (string, error)
		UpdatePayment(int64, int) error
		//UpdatePayments()
		//UpdateSubscriptions()
	}

	LoggerInterface interface {
		Debug(string, ...any)
		Error(string, ...any)
		Info(string, ...any)
		Warn(string, ...any)
	}
)
