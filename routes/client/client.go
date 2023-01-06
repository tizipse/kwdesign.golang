package client

import "github.com/gin-gonic/gin"

func Clients(routes *gin.Engine) {

	route := routes.Group("client")
	{

		RouteWeb(route)

	}

}
