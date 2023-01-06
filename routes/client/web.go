package client

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/client/web"
)

func RouteWeb(routes *gin.RouterGroup) {

	route := routes.Group("web")
	{

		categories := route.Group("categories")
		{
			categories.GET(":uri", web.ToCategoryByInformation)
		}

		banners := route.Group("banners")
		{
			banners.GET("", web.ToBanners)
		}

		picture := route.Group("picture")
		{
			picture.GET("", web.ToPicture)
		}

		setting := route.Group("setting")
		{
			setting.GET("", web.ToSetting)
		}

		contacts := route.Group("contacts")
		{
			contacts.GET("", web.ToContacts)
		}

		classifications := route.Group("classifications")
		{
			classifications.GET("", web.ToClassifications)
		}

		projects := route.Group("projects")
		{
			projects.GET("", web.ToProjectByPaginate)
			projects.GET(":id", web.ToProjectByInformation)
		}

		project := route.Group("project")
		{
			project.GET("related", web.ToProjectByRelated)
			project.GET("recommend", web.ToProjectByRecommend)
		}
	}
}
