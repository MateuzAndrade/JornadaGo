package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"mensage": "Delete Opening",
	})
}
