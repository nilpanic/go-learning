package home

import (
	"github.com/gin-gonic/gin"

	"web-tpl/app"
	"web-tpl/app/http/models"
)

func Index(ctx *gin.Context) {
	var rel []models.User

	err := app.DBW().Debug().Find(&rel).Error
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"data": rel,
	})
}
