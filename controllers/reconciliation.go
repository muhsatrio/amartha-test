package controllers

import (
	"golang-boilerplate/service"
	"golang-boilerplate/service/dto"
	"golang-boilerplate/service/request"
	"net/http"
	"strconv"
	"time"

	"encoding/csv"

	"github.com/gofiber/fiber/v2"
)

func (c Controller) reconciliation(ctx *fiber.Ctx) error {
	type transactionDto struct {
		TrxId           string    `json:"trxId"`
		Amount          int       `json:"amount"`
		Type            string    `json:"type"`
		TransactionTime time.Time `json:"transaction_time"`
	}

	type bankStatementDto struct {
		UniqueIdentifier string    `json:"unique_identifier"`
		Amount           int       `json:"amount"`
		TransactionDate  time.Time `json:"transaction_date"`
	}

	type transactionsUnmatch struct {
		Transactions   []transactionDto   `json:"transactions"`
		BankStatements []bankStatementDto `json:"bank_statements"`
	}

	type response struct {
		TransactionsProcessed int                 `json:"transactions_processed"`
		TransactionsMatched   int                 `json:"transactions_matched"`
		TransactionsUnmatch   transactionsUnmatch `json:"transactions_unmatch"`
		TotalDiscrepancies    int                 `json:"total_discrepancies"`
	}

	dateTimeFormat := "2006-01-02 15:04:05"
	dateFormat := "2006-01-02"

	// read transaction

	file, err := ctx.FormFile("transaction")
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}
	openedFile, err := file.Open()
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	transactions, err := csv.NewReader(openedFile).ReadAll()
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	// remove header

	transactions = transactions[1:]

	// read bank_statement

	file, err = ctx.FormFile("bank_statement")
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}
	openedFile, err = file.Open()
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	bankStatements, err := csv.NewReader(openedFile).ReadAll()
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	// remove header
	bankStatements = bankStatements[1:]

	var trasanctionsDto []dto.TrasanctionDto
	var bankStatementsDto []dto.BankStatementDto

	for _, trx := range transactions {
		convertedAmount, err := strconv.Atoi(trx[1])
		if err != nil {
			httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
			return ctx.Status(httpStatus).JSON(errObj)
		}

		convertedDateTime, err := time.Parse(dateTimeFormat, trx[3])
		if err != nil {
			httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
			return ctx.Status(httpStatus).JSON(errObj)
		}

		trasanctionsDto = append(trasanctionsDto, dto.TrasanctionDto{
			TrxId:           trx[0],
			Amount:          convertedAmount,
			Type:            trx[2],
			TransactionTime: convertedDateTime,
		})
	}

	for _, statement := range bankStatements {
		convertedAmount, err := strconv.Atoi(statement[1])
		if err != nil {
			httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
			return ctx.Status(httpStatus).JSON(errObj)
		}

		convertedDateTime, err := time.Parse(dateFormat, statement[2])
		if err != nil {
			httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
			return ctx.Status(httpStatus).JSON(errObj)
		}

		bankStatementsDto = append(bankStatementsDto, dto.BankStatementDto{
			UniqueIdentifier: statement[0],
			Amount:           convertedAmount,
			TransactionDate:  convertedDateTime,
		})
	}

	startDate := ctx.FormValue("start_date")

	convertedStartDate, err := time.Parse(dateFormat, startDate)
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	endDate := ctx.FormValue("end_date")

	convertedEndDate, err := time.Parse(dateFormat, endDate)
	if err != nil {
		httpStatus, errObj := errorHandler(service.InternalErrorCustom(err.Error()))
		return ctx.Status(httpStatus).JSON(errObj)
	}

	resp, serviceErr := c.ReconService.Reconciliation(request.ReconciliationRequest{
		Transactions:   trasanctionsDto,
		BankStatements: bankStatementsDto,
		StartDate:      convertedStartDate,
		EndDate:        convertedEndDate,
	})
	if serviceErr != nil {
		httpStatus, errObj := errorHandler(serviceErr)
		return ctx.Status(httpStatus).JSON(errObj)
	}

	var controllerResponse response

	controllerResponse.TotalDiscrepancies = resp.TotalDiscrepancies
	controllerResponse.TransactionsMatched = resp.TransactionsMatched
	controllerResponse.TransactionsProcessed = resp.TransactionsProcessed

	for _, data := range resp.TransactionsUnmatch.Transactions {
		controllerResponse.TransactionsUnmatch.Transactions = append(controllerResponse.TransactionsUnmatch.Transactions, transactionDto(data))
	}

	for _, data := range resp.TransactionsUnmatch.BankStatements {
		controllerResponse.TransactionsUnmatch.BankStatements = append(controllerResponse.TransactionsUnmatch.BankStatements, bankStatementDto(data))
	}

	return ctx.Status(http.StatusOK).JSON(controllerResponse)
}
