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

func ToPictures(ctx *gin.Context) {

	var settings []model.WebPicture

	app.Database.Find(&settings)

	responses := make([]wr.ToPictures, len(settings))

	for index, item := range settings {
		responses[index] = wr.ToPictures{
			Id:        item.Id,
			Label:     item.Label,
			Key:       item.Key,
			Val:       item.Val,
			Required:  item.Required,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	response.Success(ctx, responses)
}

func DoPictureByCreate(ctx *gin.Context) {

	var request web.DoPictureByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var total int64 = 0

	app.Database.Model(model.WebPicture{}).Where("`key`=?", request.Key).Count(&total)

	if total > 0 {
		response.Fail(ctx, "键已存在")
		return
	}

	picture := model.WebPicture{
		Label:    request.Label,
		Key:      request.Key,
		Val:      request.Val,
		Required: request.Required,
	}

	if cp := app.Database.Create(&picture); cp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("创建失败：%v", cp.Error))
		return
	}

	response.Success[any](ctx)
}

func DoPictureByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request web.DoPictureByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var picture model.WebPicture

	fp := app.Database.First(&picture, id)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "图片不存在")
		return
	} else if fp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("图片查找失败：%v", fp.Error))
		return
	}

	picture.Label = request.Label
	picture.Val = request.Val

	if up := app.Database.Save(&picture); up.Error != nil {
		response.Fail(ctx, fmt.Sprintf("修改失败：%v", up.Error))
		return
	}

	response.Success[any](ctx)
}

func DoPictureByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var picture model.WebPicture

	fp := app.Database.First(&picture, id)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "图片不存在")
		return
	} else if fp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("图片查找失败：%v", fp.Error))
		return
	}

	if dp := app.Database.Delete(&picture); dp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("图片删除失败：%v", dp.Error))
		return
	}

	response.Success[any](ctx)
}
