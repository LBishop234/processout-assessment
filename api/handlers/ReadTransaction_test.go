package handlers

import (
	"encoding/json"
	"main/core/bank"
	"main/core/db"
	"main/core/domain/transaction"
	"main/core/transactions"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadTransactionDetailsHandler(t *testing.T) {
	db.InitDB(true)
	bank.InitBank(bank.Naive)

	aTransaction := transaction.RndTransaction()

	err := transactions.SynchronousTransaction(aTransaction)
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
