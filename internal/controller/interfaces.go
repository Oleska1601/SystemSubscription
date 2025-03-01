package controller

import "SystemSubscription/internal/entity"

type (
	UseCaseInterface interface {
		AddPayment(*entity.SubscriptionType, int64) (string, error)
		AddSubscription(string, int64) (int64, error)
		GetLastSubscription(int64) (entity.Subscription, error)
		GetPayment(string) (entity.Payment, error)
		GetSubscriptionTypes() ([]entity.SubscriptionType, error)
		IsLastSubscription(entity.Subscription) bool
		IsPaymentActive(entity.Payment) bool
		IsSubscriptionStatusActive(entity.Subscription) bool
		LoginUser(*entity.User) (entity.User, error)
		RegisterUser(*entity.User) error
		UpdatePayment(int64, string) error
	}

	LoggerInterface interface {
		Debug(string, ...any)
		Error(string, ...any)
		Info(string, ...any)
		Warn(string, ...any)
	}
)
