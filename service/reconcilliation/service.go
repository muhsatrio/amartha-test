package reconcilliation

import (
	"golang-boilerplate/service"
	"golang-boilerplate/service/request"
	"golang-boilerplate/service/response"
)

type ReconciliationServiceInterface interface {
	Reconciliation(request request.ReconciliationRequest) (response response.ReconciliationResponse, errService service.Error)
}

type ReconciliationService struct {
}

var _ ReconciliationServiceInterface = ReconciliationService{}
