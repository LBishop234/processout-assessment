package transactions

import (
	"main/core/domain/card"
	"main/core/domain/transaction"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSynchronousMerchantTransaction(t *testing.T) {
	testReq := transaction.NewTransaction(card.RndCardNo(), time.Now().Add((365+61)*24*time.Hour), 100.0, transaction.GBP, card.RndCardCVV())

	err := SynchronousMerchantTransaction(testReq)
	require.Nil(t, err)

	if testReq.State() != transaction.Completed && testReq.State() != transaction.Declined {
		assert.Fail(t, "Transaction state should be either Completed or Declined", "State == %s", testReq.State())
	}
}
