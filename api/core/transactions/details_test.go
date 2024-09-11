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

func TestRead(t *testing.T) {
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

	require.NoError(t, SynchronousMerchantTransaction(aTransaction))
	gotTransaction, err := ReadTransaction(aTransaction.ID)
	require.NoError(t, err)
	assert.Equal(t, aTransaction, gotTransaction)
}
