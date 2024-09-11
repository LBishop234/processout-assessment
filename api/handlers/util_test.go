package handlers

import "github.com/gin-gonic/gin"

func testRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/sync", SynchronousMerchantTransactionHandler)
	router.GET("/transaction/:id", ReadTransactionDetailsHandler)
	return router
}
