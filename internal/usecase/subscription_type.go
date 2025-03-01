package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (u *Usecase) GetSubscriptionTypes() ([]entity.SubscriptionType, error) {
	ctx, _ := context.WithCancel(context.Background())
	subscriptionTypes, err := u.pgRepo.GetSubscriptionTypes(ctx)
	if err != nil {
		u.logger.Error("usecase-usecaseAPI GetSubscriptionTypes u.pgRepo.GetSubscriptionTypes", slog.Any("error", err))
		return nil, errors.New("error of getting subscription types")
	}
	fmt.Println(subscriptionTypes)
	return subscriptionTypes, nil
}
