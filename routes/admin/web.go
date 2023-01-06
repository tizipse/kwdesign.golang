package admin

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/admin/web"
)

func RouteWeb(routes *gin.RouterGroup) {

	route := routes.Group("web")
	{
		banners := route.Group("banners")
		{
			banners.GET("", web.ToBannerByPaginate)
			banners.PUT(":id", web.DoBannerByUpdate)
			banners.DELETE(":id", web.DoBannerByDelete)
		}

		banner := route.Group("banner")
		{
			banner.POST("", web.DoBannerByCreate)
			banner.PUT("enable", web.DoBannerByEnable)
		}

		projects := route.Group("projects")
		{
			projects.GET("", web.ToProjectByPaginate)
			projects.GET(":id", web.ToProjectByInformation)
			projects.PUT(":id", web.DoProjectByUpdate)
			projects.DELETE(":id", web.DoProjectByDelete)
		}

		project := route.Group("project")
		{
			project.POST("", web.DoProjectByCreate)
			project.PUT("enable", web.DoProjectByEnable)
		}

		categories := route.Group("categories")
		{
			categories.GET("", web.ToCategories)
			categories.GET(":id", web.ToCategoryByInformation)
			categories.PUT(":id", web.DoCategoryByUpdate)
			categories.DELETE(":id", web.DoCategoryByDelete)
		}

		category := route.Group("category")
		{
			category.POST("", web.DoCategoryByCreate)
			category.PUT("enable", web.DoCategoryByEnable)
			category.PUT("is_required_picture", web.DoCategoryByIsRequiredPicture)
			category.PUT("is_required_html", web.DoCategoryByIsRequiredHtml)
		}

		classifications := route.Group("classifications")
		{
			classifications.GET("", web.ToClassifications)
			classifications.GET(":id", web.ToClassificationByInformation)
			classifications.PUT(":id", web.DoClassificationByUpdate)
			classifications.DELETE(":id", web.DoClassificationByDelete)
		}

		classification := route.Group("classification")
		{
			classification.POST("", web.DoClassificationByCreate)
			classification.GET("enable", web.ToClassificationByEnable)
			classification.PUT("enable", web.DoClassificationByEnable)
		}

		contacts := route.Group("contacts")
		{
			contacts.GET("", web.ToContactByPaginate)
			contacts.PUT(":id", web.DoContactByUpdate)
			contacts.DELETE(":id", web.DoContactByDelete)
		}

		contact := route.Group("contact")
		{
			contact.POST("", web.DoContactByCreate)
			contact.PUT("enable", web.DoContactByEnable)
		}

		pictures := route.Group("pictures")
		{
			pictures.GET("", web.ToPictures)
			pictures.PUT(":id", web.DoPictureByUpdate)
			pictures.DELETE(":id", web.DoPictureByDelete)
		}

		picture := route.Group("picture")
		{
			picture.POST("", web.DoPictureByCreate)
		}

		setting := route.Group("setting")
		{
			setting.GET("", web.ToSettingByInformation)
			setting.PUT("", web.DoSettingBySave)
		}
	}
}
