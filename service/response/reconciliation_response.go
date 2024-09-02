package response

import "golang-boilerplate/service/dto"

type ReconciliationResponse struct {
	TransactionsProcessed int
	TransactionsMatched   int
	TransactionsUnmatch   TransactionsUnmatch
	TotalDiscrepancies    int
}

type TransactionsUnmatch struct {
	Transactions   []dto.TrasanctionDto
	BankStatements []dto.BankStatementDto
}
