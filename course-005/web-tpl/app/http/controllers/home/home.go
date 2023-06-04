package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code": 0,
		"msg":  "success",
		"data": "Hello Gopher!",
	})
}
