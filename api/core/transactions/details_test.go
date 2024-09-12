package transactions

import (
	"main/core/bank"
	"main/core/db"
	"main/core/domain/transaction"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	db.InitDB(true)
	bank.InitBank(bank.Naive)

	aTransaction := transaction.RndTransaction()

	require.NoError(t, SynchronousTransaction(aTransaction))

	t.Run("Read unmasked", func(t *testing.T) {
		gotTransaction, err := ReadTransaction(aTransaction.ID, false)
		require.NoError(t, err)
		assert.Equal(t, aTransaction, gotTransaction)
	})

	t.Run("Read masked", func(t *testing.T) {
		gotTransaction, err := ReadTransaction(aTransaction.ID, true)
		require.NoError(t, err)
		assert.NotEqual(t, aTransaction.CardNo, gotTransaction.CardNo)
		assert.Equal(t, aTransaction.CardNo[len(aTransaction.CardNo)-4:], gotTransaction.CardNo[len(aTransaction.CardNo)-4:])
		assert.NotEqual(t, aTransaction.CVV, gotTransaction.CVV)
	})
}
