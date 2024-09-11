package bank

import "main/domain"

type Bank interface {
	SynchronousPayment(req *domain.Transaction) error
}
