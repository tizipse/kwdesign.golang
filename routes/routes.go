package routes

import (
	"github.com/gin-gonic/gin"
	"kwd/app/middleware/basic"
	"kwd/kernel/app"
	"kwd/routes/admin"
	"kwd/routes/client"
)

func Routes(route *gin.Engine) {

	route.Use(basic.LoggerMiddleware())

	route.Static("/upload", app.Dir.Runtime+"/upload")

	admin.Admins(route)
	client.Clients(route)
}
