package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	hostPort = ":3030"
)

func main() {
	router := gin.Default()

	router.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})
	router.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})

	if err := router.Run(hostPort); err != nil {
		panic(err)
	}
}
