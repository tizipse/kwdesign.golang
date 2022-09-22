package admin

import (
	"github.com/gin-gonic/gin"
	"kwd/app/controller/admin/site"
)

func RouteSite(route *gin.RouterGroup) {

	sg := route.Group("site")
	{
		helper := sg.Group("helper")
		{
			helper.GET("apis", site.ToApiByList)
		}

		admins := sg.Group("admins")
		{
			admins.GET("", site.ToAdminByPaginate)
			admins.PUT(":id", site.DoAdminByUpdate)
			admins.DELETE(":id", site.DoAdminByDelete)
		}

		admin := sg.Group("admin")
		{
			admin.POST("", site.DoAdminByCreate)
			admin.PUT("enable", site.DoAdminByEnable)
		}

		permissions := sg.Group("permissions")
		{
			permissions.GET("", site.ToPermissionByTree)
			permissions.PUT(":id", site.DoPermissionByUpdate)
			permissions.DELETE(":id", site.DoPermissionByDelete)
		}

		permission := sg.Group("permission")
		{
			permission.GET("parents", site.ToPermissionByParents)
			permission.GET("self", site.ToPermissionBySelf)
			permission.POST("", site.DoPermissionByCreate)
		}

		roles := sg.Group("roles")
		{
			roles.GET("", site.ToRoleByPaginate)
			roles.PUT(":id", site.DoRoleByUpdate)
			roles.DELETE(":id", site.DoRoleByDelete)
		}

		role := sg.Group("role")
		{
			role.POST("", site.DoRoleByCreate)
			role.GET("enable", site.ToRoleByEnable)
		}

		modules := sg.Group("modules")
		{
			modules.GET("", site.ToModuleByList)
			modules.PUT(":id", site.DoModuleByUpdate)
			modules.DELETE(":id", site.DoModuleByDelete)
		}

		module := sg.Group("module")
		{
			module.POST("", site.DoModuleByCreate)
			module.GET("enable", site.ToModuleByEnable)
			module.PUT("enable", site.DoModuleByEnable)
		}
	}
}
