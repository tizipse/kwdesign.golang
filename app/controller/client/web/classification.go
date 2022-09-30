package web

import (
	"github.com/gin-gonic/gin"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/response/client/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToClassifications(ctx *gin.Context) {

	var classifications []model.WebClassification

	app.Database.Order("`order` asc, `id` asc").Find(&classifications, "`is_enable`=?", constant.IsEnableYes)

	responses := make([]web.ToClassifications, len(classifications))

	for index, item := range classifications {
		responses[index] = web.ToClassifications{
			Id:          item.Id,
			Name:        item.Name,
			Title:       item.Title,
			Keyword:     item.Keyword,
			Description: item.Description,
		}
	}

	response.Success(ctx, responses)
}
