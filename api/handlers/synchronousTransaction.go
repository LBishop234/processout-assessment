package handlers

import (
	"main/core/domain/transaction"
	"main/core/transactions"

	"github.com/gin-gonic/gin"
)

func synchronousTransactionHandler(c *gin.Context) {
	var aTransaction transaction.Transaction
	err := c.BindJSON(&aTransaction)
	if err != nil {
		c.Status(400)
		c.Error(err)
		return
	}

	err = transactions.SynchronousTransaction(&aTransaction)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	c.JSON(200, transaction.NewTransactionStatus(aTransaction.ID, aTransaction.State))
}
