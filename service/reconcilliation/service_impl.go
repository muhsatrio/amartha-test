package reconcilliation

import (
	"golang-boilerplate/service"
	"golang-boilerplate/service/dto"
	serviceReq "golang-boilerplate/service/request"
	serviceResp "golang-boilerplate/service/response"
)

var mapTransaction map[string]dto.TrasanctionDto
var mapBankStatement map[string]dto.BankStatementDto

// Reconciliation implements ReconciliationServiceInterface.
func (r ReconciliationService) Reconciliation(request serviceReq.ReconciliationRequest) (response serviceResp.ReconciliationResponse, errService service.Error) {

	mapTransaction = make(map[string]dto.TrasanctionDto)
	mapBankStatement = make(map[string]dto.BankStatementDto)

	if !request.StartDate.Before(request.EndDate) {
		errService = service.ErrInvalidInput
		return
	}

	for _, trx := range request.Transactions {
		mapTransaction[trx.TrxId] = trx
	}

	for _, bankState := range request.BankStatements {
		mapBankStatement[bankState.UniqueIdentifier] = bankState
	}

	for _, trx := range request.Transactions {
		if trx.TransactionTime.After(request.StartDate) && trx.TransactionTime.Before(request.EndDate) {
			response.TransactionsProcessed++
			_, isExist := mapBankStatement[trx.TrxId]
			if isExist {
				response.TransactionsMatched++
				response.TotalDiscrepancies += trx.Amount
			} else {
				response.TransactionsUnmatch.Transactions = append(response.TransactionsUnmatch.Transactions, trx)
			}
		}
	}

	for _, bankState := range request.BankStatements {
		_, isExist := mapTransaction[bankState.UniqueIdentifier]
		if !isExist {
			response.TransactionsUnmatch.BankStatements = append(response.TransactionsUnmatch.BankStatements, bankState)
		}
	}
	return
}
