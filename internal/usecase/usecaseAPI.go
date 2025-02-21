package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) GetSubscriptionTypes() ([]entity.SubscriptionType, error) {
	ctx, _ := context.WithCancel(context.Background())
	subscriptionTypes, err := u.pgRepo.GetSubscriptionTypes(ctx)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI GetSubscriptionTypes u.pgRepo.GetSubscriptionTypes", slog.Any("error", err))
		return nil, errors.New("internal server error")
	}
	return subscriptionTypes, nil
}

func (u *Usecase) GetLastSubscription(userID int64) (entity.Subscription, error) {
	ctx, _ := context.WithCancel(context.Background())
	lastSubscription, err := u.pgRepo.GetLastSubscription(ctx, userID)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI GetLastSubscription u.pgRepo.GetLastSubscription", slog.Any("error", err))
		return entity.Subscription{}, errors.New("internal server error")
	}
	return lastSubscription, nil
}

func (u *Usecase) GetNews(userID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	lastSubscription, err := u.pgRepo.GetLastSubscription(ctx, userID)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI GetLastSubscription u.pgRepo.GetLastSubscription", slog.Any("error", err))
		return errors.New("internal server error")
	}
	if lastSubscription.Status != "active" {
		return errors.New("subscription is not active")
	}
	return nil
}
