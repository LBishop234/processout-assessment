package ports

import (
	"bytes"
	"encoding/json"
	"main/core/bank"
	"main/core/db"
	"main/core/domain/transaction"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSynchronousTransactionHandler(t *testing.T) {
	db.InitDB(true)
	bank.InitBank(bank.Naive)

	aTransaction := transaction.RndTransaction()

	r := testRouter()
	w := httptest.NewRecorder()

	tBytes, err := json.Marshal(aTransaction)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/transaction/sync", bytes.NewReader(tBytes))
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var aStatus transaction.TransactionStatus
	err = json.Unmarshal(w.Body.Bytes(), &aStatus)
	require.NoError(t, err)

	assert.Equal(t, aTransaction.ID, aStatus.ID)
	assert.Contains(t, []transaction.TransactionState{transaction.Successful, transaction.Unsuccessful}, aStatus.State)
}
