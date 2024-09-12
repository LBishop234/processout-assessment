package db

import (
	"main/core/domain/card"
	"main/core/domain/transaction"

	"github.com/jmoiron/sqlx"
)

type transactionTarget struct {
	ID          string  `db:"id"`
	Timestamp   int64   `db:"timestamp"`
	CardNo      string  `db:"card_no"`
	ExpiryMonth int8    `db:"expiry_month"`
	ExpiryYear  int     `db:"expiry_year"`
	CVV         int16   `db:"cvv"`
	Currency    string  `db:"currency"`
	Amount      float64 `db:"amount"`
	State       string  `db:"state"`
}

func newTransactionTarget(t *transaction.Transaction) *transactionTarget {
	return &transactionTarget{
		ID:          t.ID,
		Timestamp:   t.UnixTimestamp,
		CardNo:      t.CardNo.Prettify(),
		ExpiryMonth: t.Expiry.Month,
		ExpiryYear:  t.Expiry.Year,
		CVV:         int16(t.CVV),
		Currency:    string(t.Currency),
		Amount:      t.Amount,
		State:       string(t.State),
	}
}

func (t *transactionTarget) toTransaction() (*transaction.Transaction, error) {
	aCardNo, err := card.NewCardNo(t.CardNo)
	if err != nil {
		return nil, err
	}

	return &transaction.Transaction{
		ID:            t.ID,
		UnixTimestamp: t.Timestamp,
		CardNo:        aCardNo,
		Expiry:        &card.CardExpiry{Month: t.ExpiryMonth, Year: t.ExpiryYear},
		CVV:           card.CardCVV(t.CVV),
		Currency:      transaction.Currency(t.Currency),
		Amount:        t.Amount,
		State:         transaction.TransactionState(t.State),
	}, nil
}

func createTransactionsTable(aDB *sqlx.DB) error {
	_, err := aDB.Exec(
		`CREATE TABLE IF NOT EXISTS transactions (
			id VARCHAR(32) PRIMARY KEY NOT NULL,
			trans_time INT NOT NULL,
			card_no VARCHAR(16) NOT NULL,
			expiry_month INT NOT NULL,
			expiry_year INT NOT NULL,
			cvv INT16 NOT NULL,
			currency VARCHAR(3) NOT NULL,
			amount DECIMAL(10, 2) NOT NULL,
			state VARCHAR(32) NOT NULL
		);`,
	)
	return err
}

func InsertTransaction(db *sqlx.DB, t *transaction.Transaction) error {
	_, err := db.NamedExec(
		`INSERT INTO transactions (id, trans_time, card_no, expiry_month, expiry_year, cvv, currency, amount, state)
		VALUES (:id, :timestamp, :card_no, :expiry_month, :expiry_year, :cvv, :currency, :amount, :state);`,
		newTransactionTarget(t),
	)
	return err
}

func SelectTransaction(aDB *sqlx.DB, id string) (*transaction.Transaction, error) {
	rows, err := aDB.Queryx(`
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

	var aTrgt transactionTarget
	err = rows.Scan(&aTrgt.ID, &aTrgt.Timestamp, &aTrgt.CardNo, &aTrgt.ExpiryMonth, &aTrgt.ExpiryYear, &aTrgt.CVV, &aTrgt.Currency, &aTrgt.Amount, &aTrgt.State)
	if err != nil {
		return nil, err
	}

	return aTrgt.toTransaction()
}
