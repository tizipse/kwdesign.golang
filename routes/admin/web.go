package admin

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/admin/web"
)

func RouteWeb(route *gin.RouterGroup) {

	wb := route.Group("web")
	{
		banners := wb.Group("banners")
		{
			banners.GET("", web.ToBannerByPaginate)
			banners.PUT(":id", web.DoBannerByUpdate)
			banners.DELETE(":id", web.DoBannerByDelete)
		}

		banner := wb.Group("banner")
		{
			banner.POST("", web.DoBannerByCreate)
			banner.PUT("enable", web.DoBannerByEnable)
		}

		settings := wb.Group("settings")
		{
			settings.GET("", web.ToSettingByInformation)
		}

		setting := wb.Group("setting")
		{
			setting.PUT("", web.DoSettingBySave)
		}
	}
}
