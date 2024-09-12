package ports

import (
	"main/core/domain/card"
	"main/core/domain/transaction"
	"time"
)

// Using an additional structure to mediate between the domain and the user interface, separating user interface and internal design concerns.
type transactionTarget struct {
	ID            string  `json:"id"`
	TimestampUnix int64   `json:"timestamp_unix"`
	CardNo        string  `json:"card_no"`
	ExpiryMonth   int8    `json:"expiry_month"`
	ExpiryYear    int     `json:"expiry_year"`
	CVV           string  `json:"cvv"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	State         string  `json:"state"`
}

func newTransactionTarget(t *transaction.Transaction) *transactionTarget {
	return &transactionTarget{
		ID:            t.ID,
		TimestampUnix: t.UnixTimestamp,
		CardNo:        t.CardNo.Prettify(),
		ExpiryMonth:   t.Expiry.Month,
		ExpiryYear:    t.Expiry.Year,
		CVV:           t.CVV.String(),
		Currency:      string(t.Currency),
		Amount:        t.Amount,
		State:         string(t.State),
	}
}

func (t *transactionTarget) parseTransaction() (*transaction.Transaction, error) {
	cardNo, err := card.NewCardNo(t.CardNo)
	if err != nil {
		return nil, err
	}

	return transaction.NewTransaction(
		time.Unix(t.TimestampUnix, 0),
		cardNo,
		card.CardExpiry{
			Month: t.ExpiryMonth,
			Year:  t.ExpiryYear,
		},
		t.Amount,
		transaction.Currency(t.Currency),
		card.CardCVV(t.CVV),
	)
}
