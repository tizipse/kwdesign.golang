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

		categories := wb.Group("categories")
		{
			categories.GET("", web.ToCategories)
			categories.GET(":id", web.ToCategoryByInformation)
			categories.PUT(":id", web.DoCategoryByUpdate)
			categories.DELETE(":id", web.DoCategoryByDelete)
		}

		category := wb.Group("category")
		{
			category.POST("", web.DoCategoryByCreate)
			category.PUT("enable", web.DoCategoryByEnable)
		}

		contacts := wb.Group("contacts")
		{
			contacts.GET("", web.ToContactByPaginate)
			contacts.PUT(":id", web.DoContactByUpdate)
			contacts.DELETE(":id", web.DoContactByDelete)
		}

		contact := wb.Group("contact")
		{
			contact.POST("", web.DoContactByCreate)
			contact.PUT("enable", web.DoContactByEnable)
		}

		setting := wb.Group("setting")
		{
			setting.GET("", web.ToSettingByInformation)
			setting.PUT("", web.DoSettingBySave)
		}
	}
}
