package ports

import (
	"main/core/domain/transaction"
	"main/core/transactions"

	"github.com/gin-gonic/gin"
)

func synchronousTransactionHandler(c *gin.Context) {
	var aTrgt transactionTarget
	err := c.BindJSON(&aTrgt)
	if err != nil {
		c.Status(400)
		c.Error(err)
		return
	}

	aTransaction, err := aTrgt.parseTransaction()
	if err != nil {
		c.Status(400)
		c.Error(err)
	}

	if err := aTransaction.Validate(); err != nil {
		c.Status(400)
		c.Error(err)
		return
	}

	err = transactions.SynchronousTransaction(aTransaction)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	c.JSON(200, transaction.NewTransactionStatus(aTransaction.ID, aTransaction.State))
}
