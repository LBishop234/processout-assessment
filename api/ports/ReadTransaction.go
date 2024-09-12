package ports

import (
	"main/core/transactions"

	"github.com/gin-gonic/gin"
)

func readTransactionHandler(c *gin.Context) {
	transactionID := c.Param("id")

	aTransaction, err := transactions.ReadTransaction(transactionID, true)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	c.JSON(200, newTransactionTarget(aTransaction))
}
