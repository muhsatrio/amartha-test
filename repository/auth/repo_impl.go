package auth

import (
	"golang-boilerplate/repository/dto"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

// GenerateJwt implements AuthRepoInterface.
func (a authRepo) GenerateJwt(payload dto.AuthDto) (token string, err error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		jwt.RegisteredClaims{
			Issuer:    payload.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(a.config.Expiry) * time.Hour)),
		},
		payload.Username,
	})

	token, err = newToken.SignedString([]byte(a.config.SigningKey))

	if err!=nil {
		log.Error("authRepo.GenerateJwt error: " + err.Error())
		return 
	}

	return
}
