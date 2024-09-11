package transactions

import (
	"main/core/db"
	"main/core/domain/transaction"
)

func ReadTransaction(id string) (*transaction.Transaction, error) {
	rows, err := db.GetDB().Queryx(`
	 	SELECT id, trans_time, card_no, expiry_month, expiry_year, cvv, currency, amount, state
	 	FROM transactions
		WHERE id = ?;`,
		id,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	rows.Next()

	var t transaction.Transaction
	if err := rows.Scan(&t.ID, &t.UnixTimestamp, &t.CardNo, &t.Expiry.Month, &t.Expiry.Year, &t.CVV, &t.Currency, &t.Amount, &t.State); err != nil {
		return nil, err
	}

	return &t, nil
}
