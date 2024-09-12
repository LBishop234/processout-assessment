package bank

import (
	"errors"
	"main/core/domain/transaction"
	"sync"
)

type BankImplementation string

const (
	Naive BankImplementation = "Naive"
)

var (
	ErrUnknownBankImplementation error = errors.New("unknown bank implementation")
)

type Bank interface {
	// Non-thread safe blocking payment method.
	SynchronousPayment(req *transaction.Transaction) error
}

var (
	// Using a global bank immutable singleton, to remove the need to handle changing bank instances whilst handling traffic.
	globalBankConnection Bank
	globalBankSingleton  sync.Once
)

func InitBank(impl BankImplementation) error {
	var err error

	globalBankSingleton.Do(func() {
		switch impl {
		case Naive:
			globalBankConnection = NewNaiveBank()
		default:
			err = ErrUnknownBankImplementation
		}
	})

	return err
}

func GetBankConnection() Bank {
	return globalBankConnection
}
