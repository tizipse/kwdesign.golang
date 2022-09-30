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

		projects := wb.Group("projects")
		{
			projects.GET("", web.ToProjectByPaginate)
			projects.GET(":id", web.ToProjectByInformation)
			projects.PUT(":id", web.DoProjectByUpdate)
			projects.DELETE(":id", web.DoProjectByDelete)
		}

		project := wb.Group("project")
		{
			project.POST("", web.DoProjectByCreate)
			project.PUT("enable", web.DoProjectByEnable)
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
			category.PUT("is_required_picture", web.DoCategoryByIsRequiredPicture)
			category.PUT("is_required_html", web.DoCategoryByIsRequiredHtml)
		}

		classifications := wb.Group("classifications")
		{
			classifications.GET("", web.ToClassifications)
			classifications.GET(":id", web.ToClassificationByInformation)
			classifications.PUT(":id", web.DoClassificationByUpdate)
			classifications.DELETE(":id", web.DoClassificationByDelete)
		}

		classification := wb.Group("classification")
		{
			classification.POST("", web.DoClassificationByCreate)
			classification.GET("enable", web.ToClassificationByEnable)
			classification.PUT("enable", web.DoClassificationByEnable)
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

		pictures := wb.Group("pictures")
		{
			pictures.GET("", web.ToPictures)
			pictures.PUT(":id", web.DoPictureByUpdate)
			pictures.DELETE(":id", web.DoPictureByDelete)
		}

		picture := wb.Group("picture")
		{
			picture.POST("", web.DoPictureByCreate)
		}

		setting := wb.Group("setting")
		{
			setting.GET("", web.ToSettingByInformation)
			setting.PUT("", web.DoSettingBySave)
		}
	}
}
