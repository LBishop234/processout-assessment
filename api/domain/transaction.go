package domain

import (
	"time"

	"github.com/google/uuid"
)

// NOTE: In a larger project I would probably partition these into relevant packages with associated
// functionality to avoid a bloated domain directory. However, for this project I have kept it simple.

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

type Transaction struct {
	id       uuid.UUID
	cardNo   []int8
	expiry   time.Time
	amount   float64
	currency Currency
	cvv      []int8
	state    TransactionState
}

func NewTransaction(cardNo []int8, expiry time.Time, amount float64, currency Currency, cvv []int8) *Transaction {
	// TODO: parameter validation

	return &Transaction{
		id:       uuid.New(),
		cardNo:   cardNo,
		expiry:   expiry,
		amount:   amount,
		currency: currency,
		cvv:      cvv,
		state:    Prior,
	}
}

func (t *Transaction) ID() uuid.UUID {
	return t.id
}

func (t *Transaction) CardNo() []int8 {
	return t.cardNo
}

func (t *Transaction) Expiry() time.Time {
	return t.expiry
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Currency() Currency {
	return t.currency
}

func (t *Transaction) CVV() []int8 {
	return t.cvv
}

func (t *Transaction) State() TransactionState {
	return t.state
}

func (t *Transaction) SetState(state TransactionState) {
	t.state = state
}
