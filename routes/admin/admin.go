package admin

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/admin/basic"
	adminMiddleware "kwd/app/middleware/admin"
	basicMiddleware "kwd/app/middleware/basic"
)

func Admins(route *gin.Engine) {
	admin := route.Group("/admin")
	admin.Use(basicMiddleware.JwtParseMiddleware())
	{
		login := admin.Group("/login")
		{
			login.POST("/account", basicMiddleware.LimitMiddleware(nil), basic.DoLoginByAccount)
			login.POST("/qrcode", basicMiddleware.LimitMiddleware(nil), basic.DoLoginByQrcode)
		}

		auth := admin.Group("")
		auth.Use(basicMiddleware.AuthMiddleware(), adminMiddleware.CasbinMiddleware())
		{
			ag := auth.Group("/account")
			{
				ag.PUT("", basic.DoAccountByUpdate)
				ag.GET("/information", basic.ToAccountByInformation)
				ag.GET("/module", basic.ToAccountByModule)
				ag.GET("/permission", basic.ToAccountByPermission)
				ag.POST("/logout", basic.DoLogout)
			}

			RouteSite(auth)
			RouteWeb(auth)

			ug := auth.Group("upload")
			{
				ug.POST("", basic.DoUploadBySimple)
			}
		}
	}
}
