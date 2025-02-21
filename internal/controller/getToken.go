package controller

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	userID  int64
	comment string
	jwt.StandardClaims
}

func (s *Server) GenerateToken(userID int64, comment string) (string, error) {
	issuedAt := time.Now().Unix()
	expiresAt := issuedAt + 3600
	claims := &Claims{
		userID:  userID,
		comment: comment,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}
