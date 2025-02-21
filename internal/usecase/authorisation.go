package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"errors"
	"log/slog"
)

func (u *Usecase) RegisterUser(user *entity.User) error {
	//try to check if user already exists
	ctx, _ := context.WithCancel(context.Background())
	checkUser, err := u.pgRepo.GetUser(ctx, user.Login)
	if err != nil {
		u.logger.Error("usecase-user LoginUser u.pgRepo.GetUser", slog.Any("error", err))
		return errors.New("internal server error")
	}
	if IsUserExists(checkUser) {
		u.logger.Error("usecase-user RegisterUser", slog.String("msg", "user already exists"))
		return errors.New("incorrect login or password")
	}
	user.PasswordSalt = GenerateSalt()
	user.PasswordHash = GenerateHash(user.Password, user.PasswordSalt)
	user.Secret = GenerateSecret()
	err = u.pgRepo.InsertUser(ctx, user)
	if err != nil {
		u.logger.Error("usecase-user RegisterUser u.pgRepo.InsertUser", slog.Any("error", err))
		return errors.New("internal server error")
	}
	return nil
}

func (u *Usecase) LoginUser(user *entity.User) (entity.User, error) {
	ctx, _ := context.WithCancel(context.Background())
	loginUser, err := u.pgRepo.GetUser(ctx, user.Login)
	if err != nil {
		u.logger.Error("usecase-user LoginUser u.pgRepo.GetUser", slog.Any("error", err))
		return entity.User{}, errors.New("internal server error")
	}
	//если пользователя с таким логином не существует или пароль неверен -> error
	if !IsUserExists(loginUser) || !IsPasswordCorrect(user.Password, loginUser) {
		u.logger.Error("usecase-user LoginUser", slog.String("msg", "user is not exists or password is incorrect"))
		return entity.User{}, errors.New("incorrect login or password")
	}
	return loginUser, nil
}
