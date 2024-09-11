package transactions

import (
	"main/core/bank"
	"main/core/domain/transaction"
	"main/core/transactions/store"
)

func init() {
	store.InitDB()
}

func SynchronousMerchantTransaction(tReq *transaction.Transaction) error {
	return bank.GetBankConnection().SynchronousPayment(tReq)
}
