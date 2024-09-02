package user

import (
	"golang-boilerplate/repository/user"
	"golang-boilerplate/service"

	"golang-boilerplate/service/dto"
)

type UserServiceInterface interface {
	Find(username string) (response dto.UserDto, errService service.Error)
	Create(user dto.UserDto) (errService service.Error)
}

type UserService struct {
	UserRepo user.UserRepoInterface
}

var _ UserServiceInterface = UserService{}
