package web

import (
	"github.com/gin-gonic/gin"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/response/client/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToContacts(ctx *gin.Context) {

	var contacts []model.WebContact

	app.Database.Order("`order` asc, `id` asc").Find(&contacts, "`is_enable`=?", constant.IsEnableYes)

	responses := make([]web.ToContacts, len(contacts))

	for index, item := range contacts {
		responses[index] = web.ToContacts{
			Id:        item.Id,
			City:      item.City,
			Address:   item.Address,
			Telephone: item.Telephone,
		}
	}

	response.Success(ctx, responses)
}
