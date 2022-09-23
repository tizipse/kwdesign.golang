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

func DoCategoryByCreate(ctx *gin.Context) {

	var request web.DoCategoryByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	category := model.WebCategory{
		Uri:         request.Uri,
		Theme:       request.Theme,
		Name:        request.Name,
		Picture:     request.Picture,
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
		Html:        request.Html,
		IsEnable:    request.IsEnable,
	}

	if cc := app.Database.Create(&category); cc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目创建失败：%v", cc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoCategoryByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request web.DoCategoryByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var category model.WebCategory

	fc := app.Database.First(&category, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "栏目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	category.Theme = request.Theme
	category.Name = request.Name
	category.Picture = request.Picture
	category.Title = request.Title
	category.Keyword = request.Keyword
	category.Description = request.Description
	category.Html = request.Html
	category.IsEnable = request.IsEnable

	if uc := app.Database.Save(&category); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目修改失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoCategoryByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var category model.WebCategory

	fc := app.Database.First(&category, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "栏目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	if dc := app.Database.Delete(&category); dc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目删除失败：%v", dc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoCategoryByEnable(ctx *gin.Context) {

	var request web.DoCategoryByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var category model.WebCategory

	fc := app.Database.First(&category, request.Id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "栏目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	category.IsEnable = request.IsEnable

	if uc := app.Database.Save(&category); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

func ToCategories(ctx *gin.Context) {

	var categories []model.WebCategory

	app.Database.Order("`id` asc").Find(&categories)

	responses := make([]wr.ToCategories, len(categories))

	for index, item := range categories {
		responses[index] = wr.ToCategories{
			Id:        item.Id,
			Theme:     item.Theme,
			Uri:       item.Uri,
			Name:      item.Name,
			Picture:   item.Picture,
			IsEnable:  item.IsEnable,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	response.Success(ctx, responses)
}

func ToCategoryByInformation(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var category model.WebCategory

	fc := app.Database.First(&category, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "栏目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	responses := wr.ToCategoryByInformation{
		Id:          category.Id,
		Theme:       category.Theme,
		Uri:         category.Uri,
		Name:        category.Name,
		Picture:     category.Picture,
		Title:       category.Title,
		Keyword:     category.Keyword,
		Description: category.Description,
		Html:        category.Html,
		IsEnable:    category.IsEnable,
	}

	response.Success(ctx, responses)

}
