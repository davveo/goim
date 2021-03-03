package main

import (
	"github.com/davveo/goim/api/controller"
	"github.com/gin-gonic/gin"
)

const (
	hostPort = ":3030"
)

func main() {
	router := gin.Default()

	router.GET("api/v1/login", controller.Login)
	router.POST("api/v1/register", controller.Register)
	router.POST("api/v1/push", controller.Push)

	if err := router.Run(hostPort); err != nil {
		panic(err)
	}
}
