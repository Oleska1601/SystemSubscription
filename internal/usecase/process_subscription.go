package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"fmt"
	"time"
)

func (u *Usecase) AddSubscription(subscriptionTypeName string, userID int64) error {
	ctx, _ := context.WithCancel(context.Background())
	subscriptionType, err := u.pgRepo.GetSubcriptionType(ctx, subscriptionTypeName)
	if err != nil {
		return fmt.Errorf("u.pgRepo.GetSubcriptionTypeID %w", err)
	}
	startDate := time.Now()
	endDate := startDate.Add(time.Duration(subscriptionType.Duration) * time.Second)
	subscription := entity.Subscription{
		StartDate:          startDate,
		EndDate:            endDate,
		SubscriptionTypeID: subscriptionType.ID,
		// //p.s по хорошему здесь делаем запрос в бд и получаем subscription_status_id где subscription_status=1
		SubscriptionStatusID: 2, //active
		UserID:               userID,
	}
	err = u.pgRepo.InsertSubscription(ctx, &subscription)
	if err != nil {
		return fmt.Errorf("u.pgRepo.InsertSubscription %w", err)
	}
	return nil
}

func (u *Usecase) IsSubscriptionStatusActive(subscription entity.Subscription) bool {
	//p.s по хорошему здесь делаем запрос в бд и получаем subscription_status_id где subscription_status=1
	return subscription.SubscriptionStatusID == 2
}

func (u *Usecase) IsLastSubscription(subscription entity.Subscription) bool {
	return subscription != entity.Subscription{}
}

func (u *Usecase) GetLastSubscription(userID int64) (entity.Subscription, error) {
	ctx, _ := context.WithCancel(context.Background())
	lastSubscription, err := u.pgRepo.GetLastSubscription(ctx, userID)
	if err != nil {
		return entity.Subscription{}, fmt.Errorf("u.pgRepo.GetLastSubscription %w", err)
	}
	return lastSubscription, nil
}
