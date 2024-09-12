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

	return db.InsertTransaction(db.GetDB(), tReq)
}
