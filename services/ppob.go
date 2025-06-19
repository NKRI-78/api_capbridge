package services

import (
	"database/sql"
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
)

func TransactionListPPOB() (map[string]any, error) {

	var transaction entities.PPOBTransactionListScan
	var dataTransaction = make([]entities.PPOBTransactionListResponse, 0)

	query := `SELECT a.id AS app_id, a.name AS app_name, i.value, i.idpel, i.product, i.created_at 
	FROM invoices i 
	INNER JOIN transactions t ON t.uid = i.transaction_id
	INNER JOIN apps a ON t.app_id = a.id
	`

	if dbPPOB == nil {
		return nil, errors.New("❌ dbPPOB connection is nil")
	}

	var rows *sql.Rows
	var err error

	rows, err = dbPPOB.Debug().Raw(query).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		errTransactionRows := dbPPOB.ScanRows(rows, &transaction)

		if errTransactionRows != nil {
			helper.Logger("error", "In Server: "+errTransactionRows.Error())
			return nil, errors.New(errTransactionRows.Error())
		}

		dataTransaction = append(dataTransaction, entities.PPOBTransactionListResponse{
			App: entities.App{
				Id:   transaction.AppId,
				Name: transaction.AppName,
			},
			Value:     transaction.Value,
			Idpel:     transaction.Idpel,
			Product:   transaction.Product,
			CreatedAt: transaction.CreatedAt,
		})
	}

	return map[string]any{
		"data": dataTransaction,
	}, nil
}
