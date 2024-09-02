package user

import (
	"golang-boilerplate/platform/mysql"
	"golang-boilerplate/repository/dto"

	"github.com/gofiber/fiber/v2/log"
)

// Find implements UserRepo.
func (repo userRepo) Find(username string) (user dto.UserDto, err error) {
	var foundUser mysql.User

	if err = repo.db.Where("username = ?", username).First(&foundUser).Error; err != nil {
		log.Error("userRepo.Find error: " + err.Error())
		return
	}

	user = dto.UserDto{
		Username: foundUser.Username,
		Password: foundUser.Password,
	}

	return
}

// Create implements UserRepo.
func (repo userRepo) Create(inputedUser dto.UserDto) (err error) {
	user := mysql.User{
		Username: inputedUser.Username,
		Password: inputedUser.Password,
	}

	if err = repo.db.Create(&user).Error; err != nil {
		log.Error("userRepo.Create error: " + err.Error())
		return
	}

	return
}
