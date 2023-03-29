package client

import (
	"github.com/gin-gonic/gin"
	"kwd/app/middleware/client"
)

func Clients(routes *gin.Engine) {

	route := routes.Group("client")
	route.Use(client.CorsMiddleware())
	{

		RouteWeb(route)

	}

}
