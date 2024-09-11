package main

import (
	"main/core/bank"
	"main/core/db"
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Package Setup
	if err := bank.InitBank(bank.Naive); err != nil {
		panic(err)
	}

	if err := db.InitDB(false); err != nil {
		panic(err)
	}

	// Router
	router := gin.New()
	router.Use(gin.Logger())

	handlers.SetupRoutes(router)

	// Server
	router.Run(":8080")
}
