package transactions

import (
	"main/core/bank"
	"main/core/domain/card"
	"main/core/domain/transaction"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	if err := bank.InitBank(bank.Naive); err != nil {
		panic(err)
	}
}

func TestSynchronousMerchantTransaction(t *testing.T) {
	aTransaction := transaction.NewTransaction(
		card.RndCardNo(),
		card.CardExpiry{
			Month: 5,
			Year:  time.Now().Year() + 2,
		},
		100.0,
		transaction.GBP,
		card.RndCardCVV(),
	)

	err := SynchronousMerchantTransaction(aTransaction)
	require.NoError(t, err)
	assert.Contains(t, []transaction.TransactionState{transaction.Completed, transaction.Declined}, aTransaction.State)
}
