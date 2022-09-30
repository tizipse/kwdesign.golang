package client

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/client/web"
)

func RouteWeb(route *gin.RouterGroup) {

	wg := route.Group("web")
	{

		categories := wg.Group("categories")
		{
			categories.GET(":uri", web.ToCategoryByInformation)
		}

		banners := wg.Group("banners")
		{
			banners.GET("", web.ToBanners)
		}

		picture := wg.Group("picture")
		{
			picture.GET("", web.ToPicture)
		}

		setting := wg.Group("setting")
		{
			setting.GET("", web.ToSetting)
		}

		contacts := wg.Group("contacts")
		{
			contacts.GET("", web.ToContacts)
		}

		classifications := wg.Group("classifications")
		{
			classifications.GET("", web.ToClassifications)
		}

		projects := wg.Group("projects")
		{
			projects.GET("", web.ToProjectByPaginate)
			projects.GET(":id", web.ToProjectByInformation)
		}

		project := wg.Group("project")
		{
			project.GET("related", web.ToProjectByRelated)
		}
	}
}
