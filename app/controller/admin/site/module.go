package site

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/request/admin/site"
	res "kwd/app/response/admin/site"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func DoModuleByCreate(ctx *gin.Context) {

	var request site.DoModuleByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var module model.SysModule

	fm := app.Database.First(&module, "`slug`=?", request.Slug)

	if fm.Error == nil {
		response.Fail(ctx, "模块已存在")
		return
	}

	module = model.SysModule{
		Slug:     request.Slug,
		Name:     request.Name,
		IsEnable: request.IsEnable,
		Order:    request.Order,
	}

	if cm := app.Database.Create(&module); cm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块创建失败：%v", cm.Error))
		return
	}

	response.Success[any](ctx)
}

func DoModuleByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request site.DoModuleByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var count int64

	app.Database.Model(model.SysModule{}).Where("`id`<>? and `slug`=?", id, request.Slug).Count(&count)

	if count > 0 {
		response.Fail(ctx, "模块已存在")
		return
	}

	var module model.SysModule

	fm := app.Database.First(&module, id)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "模块不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块查找失败：%v", fm.Error))
		return
	}

	module.Slug = request.Slug
	module.Name = request.Name
	module.IsEnable = request.IsEnable
	module.Order = request.Order

	if um := app.Database.Save(&module); um.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块修改失败：%v", um.Error))
		return
	}

	response.Success[any](ctx)
}

func DoModuleByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var module model.SysModule

	fm := app.Database.First(&module, id)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "模块不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块查找失败：%v", fm.Error))
		return
	}

	if dm := app.Database.Delete(&module); dm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块删除失败：%v", dm.Error))
		return
	}

	response.Success[any](ctx)
}

func DoModuleByEnable(ctx *gin.Context) {

	var request site.DoModuleByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var module model.SysModule

	fm := app.Database.First(&module, request.Id)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "模块不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("模块查找失败：%v", fm.Error))
		return
	}

	module.IsEnable = request.IsEnable

	if um := app.Database.Save(&module); um.Error != nil {
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", um.Error))
		return
	}

	response.Success[any](ctx)
}

func ToModuleByList(ctx *gin.Context) {

	responses := make([]res.ToModuleByList, 0)

	var modules []model.SysModule

	app.Database.Order("`order` asc").Find(&modules)

	for _, item := range modules {
		responses = append(responses, res.ToModuleByList{
			Id:        item.Id,
			Slug:      item.Slug,
			Name:      item.Name,
			IsEnable:  item.IsEnable,
			Order:     item.Order,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		})
	}

	response.Success(ctx, responses)
}

func ToModuleByEnable(ctx *gin.Context) {

	responses := make([]res.ToModuleByOnline, 0)

	var modules []model.SysModule

	app.Database.
		Where("`is_enable`=?", constant.IsEnableYes).
		Order("`order` asc").
		Find(&modules)

	for _, item := range modules {
		responses = append(responses, res.ToModuleByOnline{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	response.Success[any](ctx, responses)
}
