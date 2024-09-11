package handlers

import (
	"main/core/transactions"

	"github.com/gin-gonic/gin"
)

func ReadTransactionDetailsHandler(c *gin.Context) {
	tID := c.Param("id")

	aTransaction, err := transactions.ReadTransaction(tID)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}

	c.JSON(200, aTransaction)
}
