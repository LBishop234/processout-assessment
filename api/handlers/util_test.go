package handlers

import "github.com/gin-gonic/gin"

func testRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/sync", synchronousTransactionHandler)
	router.GET("/transaction/:id", readTransactionHandler)
	return router
}
