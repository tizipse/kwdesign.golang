package client

import "github.com/gin-gonic/gin"

func Clients(route *gin.Engine) {

	client := route.Group("client")
	{

		RouteWeb(client)

	}

}
