package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"kwd/app/model"
	"kwd/app/response/admin/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
	"kwd/kernel/validator"
)

func ToSettingByInformation(ctx *gin.Context) {

	var settings []model.WebSetting

	app.Database.Find(&settings)

	responses := make([]web.ToSettingByInformation, len(settings))

	for index, item := range settings {
		responses[index] = web.ToSettingByInformation{
			Id:        item.Id,
			Type:      item.Type,
			Label:     item.Label,
			Key:       item.Key,
			Val:       item.Val,
			Required:  item.Required,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	response.Success(ctx, responses)
}

func DoSettingBySave(ctx *gin.Context) {

	var request map[string]string

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var settings []model.WebSetting

	app.Database.Find(&settings)

	updates := make(map[int]string, 0)

	for _, item := range settings {

		req, ok := request[item.Key]

		if item.Required == model.WebSettingRequiredOfYes && (!ok || strutil.IsEmpty(req)) {
			response.FailByRequest(ctx, errors.New(item.Label+"不能为空"))
			return
		}

		if !strutil.IsEmpty(req) {

			var err error

			if item.Type == model.WebSettingTypeOfEnable {
				err = validator.Valid.Var(req, "oneof=1 2")
			} else if item.Type == model.WebSettingTypeOfUrl {
				err = validator.Valid.Var(req, "url")
			} else if item.Type == model.WebSettingTypeOfEmail {
				err = validator.Valid.Var(req, "email")
			}

			if err != nil {
				response.FailByRequest(ctx, err)
				return
			}
		}

		if req != item.Val {
			updates[item.Id] = req
		}
	}

	if len(updates) > 0 {

		tx := app.Database.Begin()

		for index, item := range updates {
			if us := tx.Model(model.WebSetting{}).Where("id=?", index).Update("val", item); us.Error != nil {
				tx.Rollback()
				response.Fail(ctx, fmt.Sprintf("保存失败：%v", us.Error))
				return
			}
		}

		tx.Commit()
	}

	response.Success[any](ctx)
}
