package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"log/slog"
	"time"
)

func (u *Usecase) GetLastSubscription(userID int64) (entity.Subscription, error) {
	ctx, _ := context.WithCancel(context.Background())
	lastSubscription, err := u.pgRepo.GetLastSubscription(ctx, userID)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI GetLastSubscription u.pgRepo.GetLastSubscription", slog.Any("error", err))
		return entity.Subscription{}, errors.New("error of getting last subscription")
	}
	return lastSubscription, nil
}

func (u *Usecase) AddSubscription(subscriptionTypeName string, userID int64) (int64, error) {
	ctx, _ := context.WithCancel(context.Background())
	subscriptionType, err := u.pgRepo.GetSubcriptionType(ctx, subscriptionTypeName)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI AddSubscription u.pgRepo.GetSubcriptionTypeID", slog.Any("error", err))
		return 0, errors.New("error of adding subscription")
	}
	startDate := time.Now()
	endDate := startDate.Add(time.Duration(subscriptionType.Duration))
	subscription := entity.Subscription{
		Status:             "active",
		StartDate:          startDate,
		EndDate:            endDate,
		SubscriptionTypeID: subscriptionType.ID,
		UserID:             userID,
	}
	subscriptionID, err := u.pgRepo.InsertSubscription(ctx, &subscription)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI AddSubscription u.pgRepo.InsertSubscription", slog.Any("error", err))
		return 0, errors.New("error of adding subscription")
	}
	return subscriptionID, nil
}

func (u *Usecase) IsSubscriptionStatusActive(subscription entity.Subscription) bool {
	return subscription.Status == "active"
}

func (u *Usecase) IsLastSubscription(subscription entity.Subscription) bool {
	return subscription != entity.Subscription{}
}
