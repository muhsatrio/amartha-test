package auth

import (
	"golang-boilerplate/platform/jwt"
	"golang-boilerplate/repository/dto"
)

type AuthRepoInterface interface {
	GenerateJwt(payload dto.AuthDto) (token string, err error)
}

type authRepo struct {
	config jwt.Config
}

func InitRepo(config jwt.Config) AuthRepoInterface {
	return authRepo{
		config: config,
	}
}

var _ AuthRepoInterface = authRepo{}
