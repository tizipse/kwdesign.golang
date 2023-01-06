package routes

import (
	"github.com/gin-gonic/gin"
	"kwd/app/middleware/basic"
	"kwd/kernel/app"
	"kwd/routes/admin"
	"kwd/routes/client"
)

func Routes(routes *gin.Engine) {

	routes.Use(basic.LoggerMiddleware())

	routes.Static("upload", app.Dir.Runtime+"/upload")

	admin.Admins(routes)
	client.Clients(routes)
}
