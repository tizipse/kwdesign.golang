package site

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kwd/app/model"
	"kwd/app/request/admin/site"
	res "kwd/app/response/admin/site"
	"kwd/kernel/api"
	"kwd/kernel/app"
	"kwd/kernel/response"
	"strings"
)

func ToApiByList(ctx *gin.Context) {

	var request site.ToApiByList

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var module model.SysModule

	fm := app.Database.First(&module, request.Module)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "模块不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块查找失败：%v", fm.Error))
		return
	}

	routes := app.Engine.Routes()

	var responses = make([]res.ToApiByList, 0)

	var permissions []model.SysPermission

	app.Database.Find(&permissions, "`module_id`=? and `method`<>? and `path`<>?", module.Id, "", "")

	var permissionsCache = make(map[string]bool, 0)

	if len(permissions) > 0 {
		for _, item := range permissions {
			permissionsCache[api.OmitKey(item.Method, item.Path)] = true
		}
	}

	for _, item := range routes {

		urls := strings.Split(item.Path, "/")

		if len(urls) >= 3 && urls[1] == "admin" && urls[2] == module.Slug {
			mark := true
			if _, exist := api.OmitOfCache[api.OmitKey(item.Method, item.Path)]; exist {
				mark = false
			}

			if mark && len(permissionsCache) > 0 {
				if _, exist := permissionsCache[api.OmitKey(item.Method, item.Path)]; exist {
					mark = false
				}
			}

			if mark {
				responses = append(responses, res.ToApiByList{
					Method: item.Method,
					Path:   item.Path,
				})
			}
		}
	}

	response.Success(ctx, responses)
}
