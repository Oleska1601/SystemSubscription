package usecase

import (
	"SystemSubscription/internal/entity"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
)

//обработка корректной аутентификации пользователя в системе

func (u *Usecase) IsUserExists(loginUser entity.User) bool {
	return loginUser != entity.User{}
}

func (u *Usecase) GenerateHash(password string, salt string) string {
	passwordHash := sha256.Sum256(append([]byte(password), []byte(salt)...))
	return base64.StdEncoding.EncodeToString(passwordHash[:])
}

func (u *Usecase) GenerateSalt() string {
	saltBytes := make([]byte, 16)
	rand.Read(saltBytes)
	return base64.StdEncoding.EncodeToString(saltBytes)
}

func (u *Usecase) IsPasswordCorrect(inputPassword string, loginUser entity.User) bool {
	inputPasswordHash := u.GenerateHash(inputPassword, loginUser.PasswordSalt)
	fmt.Println(inputPasswordHash)
	return inputPasswordHash == loginUser.PasswordHash
}

func (u *Usecase) RegisterUser(user *entity.User) error {
	//try to check if user already exists
	ctx, _ := context.WithCancel(context.Background())
	checkUser, err := u.pgRepo.GetUser(ctx, user.Login)
	if err != nil {
		u.logger.Error("usecase LoginUser u.pgRepo.GetUser", slog.Any("error", err))
		return errors.New("error with register")
	}
	if u.IsUserExists(checkUser) {
		u.logger.Error("usecase RegisterUser", slog.String("msg", "user already exists"))
		return errors.New("incorrect login or password")
	}
	user.PasswordSalt = u.GenerateSalt()
	user.PasswordHash = u.GenerateHash(user.Password, user.PasswordSalt)
	err = u.pgRepo.InsertUser(ctx, user)
	if err != nil {
		u.logger.Error("usecase RegisterUser u.pgRepo.InsertUser", slog.Any("error", err))
		return errors.New("error with register")
	}
	return nil
}

func (u *Usecase) LoginUser(user *entity.User) (entity.User, error) {
	ctx, _ := context.WithCancel(context.Background())
	loginUser, err := u.pgRepo.GetUser(ctx, user.Login)
	if err != nil {
		u.logger.Error("usecase LoginUser u.pgRepo.GetUser", slog.Any("error", err))
		return entity.User{}, errors.New("error with login")
	}
	//если пользователя с таким логином не существует или пароль неверен -> error
	if !u.IsUserExists(loginUser) || !u.IsPasswordCorrect(user.Password, loginUser) {
		fmt.Println(user.Password, loginUser, loginUser.PasswordHash)
		u.logger.Error("usecase LoginUser", slog.String("msg", "user is not exists or password is incorrect"))
		return entity.User{}, errors.New("incorrect login or password")
	}
	return loginUser, nil
}
