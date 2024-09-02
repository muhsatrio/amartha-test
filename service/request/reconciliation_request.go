package request

import (
	"golang-boilerplate/service/dto"
	"time"
)

type ReconciliationRequest struct {
	Transactions   []dto.TrasanctionDto
	BankStatements []dto.BankStatementDto
	StartDate      time.Time
	EndDate        time.Time
}
