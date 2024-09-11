package transactions

import (
	"main/core/bank"
	"main/core/db"
	"main/core/domain/card"
	"main/core/domain/transaction"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSynchronousMerchantTransaction(t *testing.T) {
	db.InitDB(true)
	bank.InitBank(bank.Naive)

	aTransaction := transaction.NewTransaction(
		time.Now(),
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
