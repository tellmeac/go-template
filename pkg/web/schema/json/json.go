package json

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func APIOK(ctx *gin.Context, v any) error {
	return APIOKWithStatus(ctx, http.StatusOK, v)
}

func APIOKWithStatus(ctx *gin.Context, status int, v any) error {
	ctx.JSON(status, v)
	return nil
}

func APIError(ctx *gin.Context, status int, err error) error {
	ctx.JSON(status, gin.H{
		"message": err.Error(),
	})

	return nil
}
