package dto

import "time"

type BankStatementDto struct {
	UniqueIdentifier string
	Amount           int
	TransactionDate  time.Time
}
