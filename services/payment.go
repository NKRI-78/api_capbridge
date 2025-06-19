package services

import (
	"database/sql"
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
)

func TransactionListPayment(orderId string) (map[string]any, error) {

	var transaction entities.PaymentTransactionListScan
	var dataTransaction = make([]entities.PaymentTransactionListResponse, 0)

	query := `
		SELECT 
			orderId AS order_id,
			app,
			grossAmount AS gross_amount,
			totalAmount AS total_amount,
			transactionStatus AS transaction_status,
			createdAt AS created_at
		FROM 
			Payments
		WHERE 
			orderId LIKE ?
		ORDER BY 
			createdAt DESC
	`

	if dbPayment == nil {
		return nil, errors.New("❌ dbPayment connection is nil")
	}

	var rows *sql.Rows
	var err error

	likePattern := orderId + "%"

	rows, err = dbPayment.Debug().Raw(query, likePattern).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		errTransactionRows := dbPayment.ScanRows(rows, &transaction)

		if errTransactionRows != nil {
			helper.Logger("error", "In Server: "+errTransactionRows.Error())
			return nil, errors.New(errTransactionRows.Error())
		}

		dataTransaction = append(dataTransaction, entities.PaymentTransactionListResponse{
			App:               transaction.App,
			OrderId:           transaction.OrderId,
			GrossAmount:       transaction.GrossAmount,
			TotalAmount:       transaction.TotalAmount,
			TransactionStatus: transaction.TransactionStatus,
			CreatedAt:         transaction.CreatedAt,
		})
	}

	return map[string]any{
		"data": dataTransaction,
	}, nil
}
