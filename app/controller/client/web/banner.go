package web

import (
	"github.com/gin-gonic/gin"
	"kwd/app/constant"
	"kwd/app/model"
	wr "kwd/app/response/client/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToBanners(ctx *gin.Context) {

	var banners []model.WebBanner

	app.Database.Order("`order` asc, `id` asc").Find(&banners, "`is_enable`=?", constant.IsEnableYes)

	responses := make([]wr.ToBanners, len(banners))

	for index, item := range banners {
		responses[index] = wr.ToBanners{
			Id:      item.Id,
			Theme:   item.Theme,
			Picture: item.Picture,
			Name:    item.Name,
			Target:  item.Target,
			Url:     item.Url,
		}
	}

	response.Success(ctx, responses)

}
