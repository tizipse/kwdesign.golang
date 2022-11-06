package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/request/admin/web"
	wr "kwd/app/response/admin/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func DoClassificationByCreate(ctx *gin.Context) {

	var request web.DoClassificationByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	classification := model.WebClassification{
		Id:          app.Snowflake.Generate().String(),
		Name:        request.Name,
		Alias:       request.Alias,
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
		Order:       request.Order,
		IsEnable:    request.IsEnable,
	}

	if cc := app.Database.Create(&classification); cc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("分类创建失败：%v", cc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoClassificationByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request web.DoClassificationByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var classification model.WebClassification

	fc := app.Database.First(&classification, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "分类不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("分类查找失败：%v", fc.Error))
		return
	}

	classification.Name = request.Name
	classification.Alias = request.Alias
	classification.Title = request.Title
	classification.Keyword = request.Keyword
	classification.Description = request.Description
	classification.Order = request.Order
	classification.IsEnable = request.IsEnable

	if uc := app.Database.Save(&classification); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("分类修改失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoClassificationByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var classification model.WebClassification

	fc := app.Database.First(&classification, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "分类不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	if dc := app.Database.Delete(&classification); dc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("分类删除失败：%v", dc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoClassificationByEnable(ctx *gin.Context) {

	var request web.DoClassificationByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var classification model.WebClassification

	fc := app.Database.First(&classification, request.Id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "分类不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("分类查找失败：%v", fc.Error))
		return
	}

	classification.IsEnable = request.IsEnable

	if uc := app.Database.Save(&classification); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

func ToClassifications(ctx *gin.Context) {

	var classifications []model.WebClassification

	app.Database.
		Order("`order` asc, `id` asc").
		Find(&classifications)

	responses := make([]wr.ToClassifications, len(classifications))

	for index, item := range classifications {
		responses[index] = wr.ToClassifications{
			Id:        item.Id,
			Name:      item.Name,
			Alias:     item.Alias,
			Order:     item.Order,
			IsEnable:  item.IsEnable,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	response.Success(ctx, responses)
}

func ToClassificationByEnable(ctx *gin.Context) {

	var classifications []model.WebClassification

	app.Database.Order("`order` asc, `id` asc").Find(&classifications, "`is_enable`=?", constant.IsEnableYes)

	responses := make([]wr.ToClassificationByEnable, len(classifications))

	for index, item := range classifications {
		responses[index] = wr.ToClassificationByEnable{
			Id:   item.Id,
			Name: item.Name,
		}
	}

	response.Success(ctx, responses)
}

func ToClassificationByInformation(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var classification model.WebClassification

	fc := app.Database.First(&classification, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "栏目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	responses := wr.ToClassificationByInformation{
		Id:          classification.Id,
		Name:        classification.Name,
		Alias:       classification.Alias,
		Title:       classification.Title,
		Keyword:     classification.Keyword,
		Description: classification.Description,
		Order:       classification.Order,
		IsEnable:    classification.IsEnable,
	}

	response.Success(ctx, responses)

}
