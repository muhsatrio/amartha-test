package dto

import "time"

type TrasanctionDto struct {
	TrxId           string
	Amount          int
	Type            string
	TransactionTime time.Time
}
