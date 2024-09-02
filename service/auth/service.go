package auth

import (
	"golang-boilerplate/repository/auth"
	"golang-boilerplate/service"
	"golang-boilerplate/service/request"
	"golang-boilerplate/service/response"
	userservice "golang-boilerplate/service/user"
)

type AuthServiceInterface interface {
	Login(request request.AuthRequest) (response response.AuthResponse, errService service.Error)
}

type AuthService struct {
	UserService userservice.UserServiceInterface
	AuthRepo    auth.AuthRepoInterface
}

var _ AuthServiceInterface = AuthService{}
