package transaction

import (
	"main/core/domain/card"

	"github.com/google/uuid"
)

type (
	Currency         string
	TransactionState string
)

const (
	GBP Currency = "GBP"
	USD Currency = "USD"
	EUR Currency = "EUR"

	Prior     TransactionState = "Prior"
	Pending   TransactionState = "Pending"
	Approved  TransactionState = "Approved"
	Completed TransactionState = "Completed"
	Declined  TransactionState = "Declined"
)

// Transaction represents a transaction entity.
// Must be passed by reference.
type Transaction struct {
	ID       uuid.UUID        `json:"id"`
	CardNo   card.CardNo      `json:"card_no"`
	Expiry   card.CardExpiry  `json:"expiry"`
	Amount   float64          `json:"amount"`
	Currency Currency         `json:"currency"`
	CVV      card.CardCVV     `json:"cvv"`
	State    TransactionState `json:"state"`
}

func NewTransaction(cardNo card.CardNo, expiry card.CardExpiry, amount float64, currency Currency, cvv card.CardCVV) *Transaction {
	// TODO: parameter validation

	return &Transaction{
		ID:       uuid.New(),
		CardNo:   cardNo,
		Expiry:   expiry,
		Amount:   amount,
		Currency: currency,
		CVV:      cvv,
		State:    Prior,
	}
}

type TransactionStatus struct {
	ID    uuid.UUID        `json:"id"`
	State TransactionState `json:"state"`
}

func NewTransactionStatus(id uuid.UUID, state TransactionState) *TransactionStatus {
	return &TransactionStatus{
		ID:    id,
		State: state,
	}
}
