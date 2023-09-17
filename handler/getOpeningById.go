package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpeningByIdHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opening",
	})
}
