package web

import (
	"github.com/gin-gonic/gin"
	"kwd/app/model"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToPicture(ctx *gin.Context) {

	var pictures []model.WebPicture

	app.Database.Find(&pictures)

	responses := make(map[string]string, 0)

	for _, item := range pictures {
		responses[item.Key] = item.Val
	}

	response.Success(ctx, responses)
}
