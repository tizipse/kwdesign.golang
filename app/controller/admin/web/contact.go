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

func DoContactByCreate(ctx *gin.Context) {

	var request web.DoContactByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	contact := model.WebContact{
		City:      request.City,
		Address:   request.Address,
		Telephone: request.Telephone,
		Order:     request.Order,
		IsEnable:  request.IsEnable,
	}

	if cc := app.Database.Create(&contact); cc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("联系创建失败：%v", cc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoContactByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request web.DoContactByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var contact model.WebContact

	fc := app.Database.First(&contact, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "联系不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("联系查找失败：%v", fc.Error))
		return
	}

	contact.City = request.City
	contact.Address = request.Address
	contact.Telephone = request.Telephone
	contact.IsEnable = request.IsEnable
	contact.Order = request.Order

	if uc := app.Database.Save(&contact); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("联系修改失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoContactByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var contact model.WebContact

	fc := app.Database.First(&contact, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "联系不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("联系查找失败：%v", fc.Error))
		return
	}

	if dc := app.Database.Delete(&contact); dc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("联系删除失败：%v", dc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoContactByEnable(ctx *gin.Context) {

	var request web.DoContactByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var contact model.WebContact

	fc := app.Database.First(&contact, request.Id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "联系不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("联系查找失败：%v", fc.Error))
		return
	}

	contact.IsEnable = request.IsEnable

	if uc := app.Database.Save(&contact); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

func ToContactByPaginate(ctx *gin.Context) {

	var request web.ToContactByPaginate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	responses := response.Paginate[wr.ToContactByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := app.Database.Model(model.WebContact{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var contacts []model.WebContact

		tx.
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`order` asc, `id` asc").
			Find(&contacts)

		responses.Data = make([]wr.ToContactByPaginate, len(contacts))

		for index, item := range contacts {
			responses.Data[index] = wr.ToContactByPaginate{
				Id:        item.Id,
				City:      item.City,
				Address:   item.Address,
				Telephone: item.Telephone,
				IsEnable:  item.IsEnable,
				Order:     item.Order,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	response.Success(ctx, responses)
}
