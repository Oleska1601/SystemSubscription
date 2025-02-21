package usecase

import (
	"SystemSubscription/internal/entity"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateHash(password string, salt string) string {
	passwordHash := sha256.Sum256(append([]byte(password), []byte(salt)...))
	return base64.StdEncoding.EncodeToString(passwordHash[:])
}

func GenerateSalt() string {
	saltBytes := make([]byte, 16)
	rand.Read(saltBytes)
	return base64.StdEncoding.EncodeToString(saltBytes)
}

func GenerateSecret() string {
	secretBytes := make([]byte, 32)
	rand.Read(secretBytes)
	return base64.StdEncoding.EncodeToString(secretBytes)
}

func IsUserExists(loginUser entity.User) bool {
	return loginUser != entity.User{}
}

func IsPasswordCorrect(inputPassword string, loginUser entity.User) bool {
	inputPasswordHash := GenerateHash(inputPassword, loginUser.PasswordSalt)
	return inputPasswordHash == loginUser.PasswordHash
}

// ???
/*
func (u *Usecase) IsSubscriptionExists(subscription entity.Subscription) bool {
	return subscription != entity.Subscription{}
}
*/
