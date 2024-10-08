package ports

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/transaction/:id", readTransactionHandler)
	router.POST("/transaction/sync", synchronousTransactionHandler)
}
