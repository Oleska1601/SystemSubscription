package usecase

import (
	"SystemSubscription/internal/entity"
	"errors"
	"log/slog"
)

func (u *Usecase) GetNews(userID int64) (entity.News, error) {
	lastSubscription, err := u.GetLastSubscription(userID)
	if err != nil {
		u.logger.Error("usecase GetNews u.GetLastSubscription", slog.Any("error", err))
		return entity.News{}, errors.New("error of getting news")
	}
	if !u.IsSubscriptionStatusActive(lastSubscription) {
		u.logger.Error("usecase GetNews u.IsSubscriptionStatusActive", slog.String("msg", "current subscription is not active"))
		return entity.News{}, errors.New("error of getting news")
	}
	news := entity.News{Message: "news"}
	return news, nil
}
