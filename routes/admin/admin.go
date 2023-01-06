package admin

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/admin/basic"
	adminMiddleware "kwd/app/middleware/admin"
	basicMiddleware "kwd/app/middleware/basic"
)

func Admins(routes *gin.Engine) {

	route := routes.Group("admin")
	route.Use(basicMiddleware.JwtParseMiddleware())
	{
		login := route.Group("login")
		{
			login.POST("account", basicMiddleware.LimitMiddleware(nil), basic.DoLoginByAccount)
			login.POST("qrcode", basicMiddleware.LimitMiddleware(nil), basic.DoLoginByQrcode)
		}

		auth := route.Group("")
		auth.Use(basicMiddleware.AuthMiddleware(), adminMiddleware.CasbinMiddleware())
		{
			account := auth.Group("account")
			{
				account.PUT("", basic.DoAccountByUpdate)
				account.GET("information", basic.ToAccountByInformation)
				account.GET("module", basic.ToAccountByModule)
				account.GET("permission", basic.ToAccountByPermission)
				account.POST("logout", basic.DoLogout)
			}

			RouteSite(auth)
			RouteWeb(auth)

			upload := auth.Group("upload")
			{
				upload.POST("", basic.DoUploadBySimple)
			}
		}
	}
}
