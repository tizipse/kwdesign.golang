package web

import (
	"github.com/gin-gonic/gin"
	"kwd/app/model"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToSetting(ctx *gin.Context) {

	var settings []model.WebSetting

	app.Database.Find(&settings)

	responses := make(map[string]string, 0)

	for _, item := range settings {
		responses[item.Key] = item.Val
	}

	response.Success(ctx, responses)
}
