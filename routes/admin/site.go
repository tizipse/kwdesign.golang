package admin

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/admin/site"
)

func RouteSite(routes *gin.RouterGroup) {

	route := routes.Group("site")
	{
		helper := route.Group("helper")
		{
			helper.GET("apis", site.ToApiByList)
		}

		admins := route.Group("admins")
		{
			admins.GET("", site.ToAdminByPaginate)
			admins.PUT(":id", site.DoAdminByUpdate)
			admins.DELETE(":id", site.DoAdminByDelete)
		}

		admin := route.Group("admin")
		{
			admin.POST("", site.DoAdminByCreate)
			admin.PUT("enable", site.DoAdminByEnable)
		}

		permissions := route.Group("permissions")
		{
			permissions.GET("", site.ToPermissionByTree)
			permissions.PUT(":id", site.DoPermissionByUpdate)
			permissions.DELETE(":id", site.DoPermissionByDelete)
		}

		permission := route.Group("permission")
		{
			permission.GET("parents", site.ToPermissionByParents)
			permission.GET("self", site.ToPermissionBySelf)
			permission.POST("", site.DoPermissionByCreate)
		}

		roles := route.Group("roles")
		{
			roles.GET("", site.ToRoleByPaginate)
			roles.PUT(":id", site.DoRoleByUpdate)
			roles.DELETE(":id", site.DoRoleByDelete)
		}

		role := route.Group("role")
		{
			role.POST("", site.DoRoleByCreate)
			role.GET("enable", site.ToRoleByEnable)
		}

		modules := route.Group("modules")
		{
			modules.GET("", site.ToModuleByList)
			modules.PUT(":id", site.DoModuleByUpdate)
			modules.DELETE(":id", site.DoModuleByDelete)
		}

		module := route.Group("module")
		{
			module.POST("", site.DoModuleByCreate)
			module.GET("enable", site.ToModuleByEnable)
			module.PUT("enable", site.DoModuleByEnable)
		}
	}
}
