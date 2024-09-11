package handlers

import (
	"bytes"
	"encoding/json"
	"main/core/bank"
	"main/core/db"
	"main/core/domain/card"
	"main/core/domain/transaction"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	if err := bank.InitBank(bank.Naive); err != nil {
		panic(err)
	}
}

func testRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/sync", SynchronousMerchantTransactionHandler)
	return router
}

func TestSynchronousTransactionHandler(t *testing.T) {
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

	r := testRouter()
	w := httptest.NewRecorder()

	tBytes, err := json.Marshal(aTransaction)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/sync", bytes.NewReader(tBytes))
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var aStatus transaction.TransactionStatus
	err = json.Unmarshal(w.Body.Bytes(), &aStatus)
	require.NoError(t, err)

	assert.Equal(t, aTransaction.ID, aStatus.ID)
	assert.Contains(t, []transaction.TransactionState{transaction.Completed, transaction.Declined}, aStatus.State)
}
