package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Push(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
