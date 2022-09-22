package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/model"
	"kwd/app/request/admin/web"
	wr "kwd/app/response/admin/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func DoBannerByCreate(ctx *gin.Context) {

	var request web.DoBannerByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	banner := model.WebBanner{
		Layout:   request.Layout,
		Picture:  request.Picture,
		Name:     request.Name,
		Target:   request.Target,
		Url:      request.Url,
		IsEnable: request.IsEnable,
		Order:    request.Order,
	}

	if cm := app.Database.Create(&banner); cm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("轮播创建失败：%v", cm.Error))
		return
	}

	response.Success[any](ctx)
}

func DoBannerByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request web.DoBannerByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var banner model.WebBanner

	fm := app.Database.First(&banner, id)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "轮播不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("轮播查找失败：%v", fm.Error))
		return
	}

	banner.Layout = request.Layout
	banner.Picture = request.Picture
	banner.Name = request.Name
	banner.Target = request.Target
	banner.Url = request.Url
	banner.IsEnable = request.IsEnable
	banner.Order = request.Order

	if um := app.Database.Save(&banner); um.Error != nil {
		response.Fail(ctx, fmt.Sprintf("轮播修改失败：%v", um.Error))
		return
	}

	response.Success[any](ctx)
}

func DoBannerByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var banner model.WebBanner

	fm := app.Database.First(&banner, id)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "轮播不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("轮播查找失败：%v", fm.Error))
		return
	}

	if dm := app.Database.Delete(&banner); dm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("轮播删除失败：%v", dm.Error))
		return
	}

	response.Success[any](ctx)
}

func DoBannerByEnable(ctx *gin.Context) {

	var request web.DoBannerByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var banner model.WebBanner

	fm := app.Database.First(&banner, request.Id)

	if errors.Is(fm.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "轮播不存在")
		return
	} else if fm.Error != nil {
		response.Fail(ctx, fmt.Sprintf("轮播查找失败：%v", fm.Error))
		return
	}

	banner.IsEnable = request.IsEnable

	if um := app.Database.Save(&banner); um.Error != nil {
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", um.Error))
		return
	}

	response.Success[any](ctx)
}

func ToBannerByPaginate(ctx *gin.Context) {

	var request web.ToBannerByPaginate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	responses := response.Paginate[wr.ToBannerByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := app.Database.Model(model.WebBanner{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var banners []model.WebBanner

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` desc").
			Find(&banners)

		responses.Data = make([]wr.ToBannerByPaginate, len(banners))

		for index, item := range banners {
			responses.Data[index] = wr.ToBannerByPaginate{
				Id:        item.Id,
				Layout:    item.Layout,
				Picture:   item.Picture,
				Name:      item.Name,
				Target:    item.Target,
				Url:       item.Url,
				IsEnable:  item.IsEnable,
				Order:     item.Order,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	response.Success(ctx, responses)
}
