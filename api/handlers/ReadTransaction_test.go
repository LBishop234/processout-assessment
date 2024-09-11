package handlers

import (
	"encoding/json"
	"main/core/bank"
	"main/core/db"
	"main/core/domain/card"
	"main/core/domain/transaction"
	"main/core/transactions"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadTransactionDetailsHandler(t *testing.T) {
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

	err := transactions.SynchronousMerchantTransaction(aTransaction)
	require.NoError(t, err)

	r := testRouter()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", path.Join("/transaction/", aTransaction.ID), nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var gotTransaction transaction.Transaction
	err = json.Unmarshal(w.Body.Bytes(), &gotTransaction)
	require.NoError(t, err)

	assert.Equal(t, *aTransaction, gotTransaction)
}
