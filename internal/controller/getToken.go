package controller

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	jwt.StandardClaims
}

func (s *Server) GenerateToken(userID int64, comment string) (string, error) {
	issuedAt := time.Now().Unix()
	expiresAt := issuedAt + 3600

	claims := &Claims{
		UserID:  userID,
		Comment: comment,
		StandardClaims: jwt.StandardClaims{

			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
	}
	/*

		claims := jwt.MapClaims {
			"userID": userID,
			"comment": comment,
			"issuedAt": issuedAt,
			"expiretAt": expiresAt,
		}
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}
