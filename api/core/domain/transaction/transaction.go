package transaction

import (
	"errors"
	"main/core/domain/card"
	"math/rand"
	"time"

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

	Prior        TransactionState = "Prior"
	Pending      TransactionState = "Pending"
	Approved     TransactionState = "Approved"
	Successful   TransactionState = "Successful"
	Unsuccessful TransactionState = "Unsuccessful"
)

var (
	ErrInvalidTransactionTimestamp = errors.New("invalid transaction timestamp")
	ErrInvalidTransactionAmount    = errors.New("invalid transaction amount")
)

// Transaction represents a transaction entity.
// Must be passed by reference.
type Transaction struct {
	ID            string           `json:"id"`
	UnixTimestamp int64            `json:"timestamp_unix"`
	CardNo        card.CardNo      `json:"card_no"`
	Expiry        *card.CardExpiry `json:"expiry"`
	Amount        float64          `json:"amount"`
	Currency      Currency         `json:"currency"`
	CVV           card.CardCVV     `json:"cvv"`
	State         TransactionState `json:"state"`
}

func NewTransaction(timestamp time.Time, cardNo card.CardNo, expiry card.CardExpiry, amount float64, currency Currency, cvv card.CardCVV) (*Transaction, error) {
	aTransaction := &Transaction{
		ID:            uuid.New().String(),
		UnixTimestamp: timestamp.Unix(),
		CardNo:        cardNo,
		Expiry:        &expiry,
		Amount:        amount,
		Currency:      currency,
		CVV:           cvv,
		State:         Prior,
	}

	return aTransaction, aTransaction.Validate()
}

func BlankTransaction() *Transaction {
	return &Transaction{
		Expiry: &card.CardExpiry{},
	}
}

func RndTransaction() *Transaction {
	aTransaction, err := NewTransaction(
		time.Now().Add(time.Duration(-rand.Intn(30)*int(time.Minute))),
		card.RndCardNo(),
		card.CardExpiry{
			Month: int8(rand.Intn(12)) + 1,
			Year:  time.Now().Year() + rand.Intn(5),
		},
		rand.Float64()*1000,
		GBP,
		card.RndCardCVV(),
	)
	if err != nil {
		// This should never happen
		panic(err)
	}

	return aTransaction
}

func (t *Transaction) Validate() error {
	if err := t.CardNo.Validate(); err != nil {
		return err
	}

	if err := t.CVV.Validate(); err != nil {
		return err
	}

	if t.Amount < 0 {
		return ErrInvalidTransactionAmount
	}

	if t.UnixTimestamp > time.Now().Unix() {
		return ErrInvalidTransactionTimestamp
	}

	return nil
}

func (t *Transaction) MaskDetails() {
	t.CardNo = t.CardNo.Mask()
	t.CVV = t.CVV.Mask()
}

type TransactionStatus struct {
	ID    string           `json:"id"`
	State TransactionState `json:"state"`
}

func NewTransactionStatus(id string, state TransactionState) *TransactionStatus {
	return &TransactionStatus{
		ID:    id,
		State: state,
	}
}
