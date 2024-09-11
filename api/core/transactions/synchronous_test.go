package transactions

import (
	"main/core/bank"
	"main/core/db"
	"main/core/domain/transaction"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSynchronousTransaction(t *testing.T) {
	db.InitDB(true)
	bank.InitBank(bank.Naive)

	aTransaction := transaction.RndTransaction()

	err := SynchronousTransaction(aTransaction)
	require.NoError(t, err)
	assert.Contains(t, []transaction.TransactionState{transaction.Completed, transaction.Declined}, aTransaction.State)
}
