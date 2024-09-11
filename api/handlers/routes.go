package handlers

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/transaction/:id", readTransactionHandler)
	router.POST("/transactions/sync", synchronousTransactionHandler)
}
