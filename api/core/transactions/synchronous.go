package transactions

import (
	"main/core/bank"
	"main/core/db"
	"main/core/domain/transaction"
)

func SynchronousTransaction(tReq *transaction.Transaction) error {
	if err := tReq.Validate(); err != nil {
		return err
	}

	if err := bank.GetBankConnection().SynchronousPayment(tReq); err != nil {
		return err
	}

	if _, err := db.GetDB().Exec(`
		INSERT INTO transactions (id, trans_time, card_no, expiry_month, expiry_year, cvv, currency, amount, state)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		tReq.ID, tReq.UnixTimestamp, tReq.CardNo, tReq.Expiry.Month, tReq.Expiry.Year, tReq.CVV, tReq.Currency, tReq.Amount, tReq.State,
	); err != nil {
		return err
	}

	return nil
}
