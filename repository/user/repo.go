package user

import (
	"golang-boilerplate/repository/dto"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Find(username string) (dto.UserDto, error)
	Create(dto.UserDto) error
}

type userRepo struct {
	db *gorm.DB
}

func InitRepo(db *gorm.DB) UserRepoInterface {
	return userRepo{
		db: db,
	}
}

var _ UserRepoInterface = userRepo{}
