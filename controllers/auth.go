package controllers

import (
	"golang-boilerplate/service"
	serviceRequest "golang-boilerplate/service/request"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (c Controller) login(ctx *fiber.Ctx) error {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type response struct {
		Token string `json:"token"`
	}

	body := new(request)

	if err := ctx.BodyParser(body); err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	resp, err := c.AuthService.Login(serviceRequest.AuthRequest{
		Username: body.Username,
		Password: body.Password,
	})

	if err != nil {
		httpStatus, errObj := errorHandler(err)
		return ctx.Status(httpStatus).JSON(errObj)
	}

	return ctx.Status(http.StatusOK).JSON(response{
		Token: resp.Token,
	})
}
