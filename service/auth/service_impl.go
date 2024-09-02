package auth

import (
	dtorepo "golang-boilerplate/repository/dto"
	"golang-boilerplate/service"
	"golang-boilerplate/service/request"
	"golang-boilerplate/service/response"

	"github.com/gofiber/fiber/v2/log"
)

func (svc AuthService) Login(request request.AuthRequest) (resp response.AuthResponse, errService service.Error) {
	user, err := svc.UserService.Find(request.Username)
	if err != nil {
		if err == service.ErrDataNotFound {
			errService = service.ErrUnauthorized
			return
		}
	}
	token, err := svc.AuthRepo.GenerateJwt(dtorepo.AuthDto{
		Username: user.Username,
	})
	if err != nil {
		log.Error("AuthService.Login error: " + err.Error())
		errService = service.InternalErrorCustom(err.Error())
		return
	}

	resp = response.AuthResponse{
		Token: token,
	}

	return
}
