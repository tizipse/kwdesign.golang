package web

import (
	"github.com/gin-gonic/gin"
	"kwd/app/model"
	"kwd/app/response/admin/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToSettingByInformation(ctx *gin.Context) {

	var settings []model.WebSetting

	app.Database.Find(&settings)

	responses := make([]web.ToSettingByInformation, len(settings))

	for index, item := range settings {
		responses[index] = web.ToSettingByInformation{
			Id:        item.Id,
			Type:      item.Type,
			Label:     item.Label,
			Key:       item.Key,
			Val:       item.Val,
			Required:  item.Required,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	response.Success(ctx, responses)
}

func DoSettingBySave(ctx *gin.Context) {

}
