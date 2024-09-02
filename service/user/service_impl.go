package user

import (
	dtoRepo "golang-boilerplate/repository/dto"
	"golang-boilerplate/service"
	"golang-boilerplate/service/dto"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Create implements ServiceInterface.
func (s UserService) Create(user dto.UserDto) (errService service.Error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("UserService.Create error: %s", err.Error())
		errService = service.InternalErrorCustom(err.Error())
		return
	}

	if err = s.UserRepo.Create(dtoRepo.UserDto{
		Username: user.Username,
		Password: string(bytes),
	}); err != nil {
		log.Errorf("UserService.Create error: %s", err.Error())
		errService = service.InternalErrorCustom(err.Error())
		return
	}

	return
}

// Find implements ServiceInterface.
func (s UserService) Find(username string) (response dto.UserDto, errService service.Error) {
	foundUser, err := s.UserRepo.Find(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errService = service.ErrDataNotFound
			return
		}
		log.Error("UserService.Find error: " + err.Error())
		errService = service.InternalErrorCustom(err.Error())
		return
	}
	response = dto.UserDto{
		Username: foundUser.Username,
		Password: foundUser.Password,
	}
	return
}
