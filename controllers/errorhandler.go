package controllers

import (
	"golang-boilerplate/service"
	"net/http"
)

type ErrorObj struct {
	Message string `json:"message"`
}

func errorHandler(err service.Error) (int, ErrorObj) {
	var errCode int

	switch err {
	case service.ErrInvalidInput, service.ErrRequiredFieldEmpty, service.ErrDuplicateDataAdd:
		{
			errCode = http.StatusBadRequest
		}
	case service.ErrUnauthorized:
		{
			errCode = http.StatusUnauthorized
		}
	case service.ErrForbiddenAccess:
		{
			errCode = http.StatusForbidden
		}
	case service.ErrDataNotFound:
		{
			errCode = http.StatusNotFound
		}
	default:
		{
			errCode = http.StatusInternalServerError
		}
	}

	errObj := ErrorObj{
		Message: err.Error(),
	}

	return errCode, errObj
}
