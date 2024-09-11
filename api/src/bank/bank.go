package bank

import "main/src/transactions"

type Bank interface {
	SynchronousPayment(req *transactions.Transaction) error
}
