package transactions

import (
	"main/core/bank"
	"main/core/domain/transaction"
)

func SynchronousMerchantTransaction(tReq *transaction.Transaction) error {
	return bank.GetBankConnection().SynchronousPayment(tReq)
}
