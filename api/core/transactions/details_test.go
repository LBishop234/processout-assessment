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
	gotTransaction, err := ReadTransaction(aTransaction.ID)
	require.NoError(t, err)
	assert.Equal(t, aTransaction, gotTransaction)
}
