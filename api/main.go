package main

import (
	"main/core/bank"

	"github.com/gin-gonic/gin"
)

func main() {
	// Package Setup
	err := bank.InitBank(bank.Naive)
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
}
