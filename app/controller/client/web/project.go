package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/constant"
	"kwd/app/model"
	"kwd/app/request/client/web"
	wr "kwd/app/response/client/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToProjectByPaginate(ctx *gin.Context) {

	var request web.ToProjectByPaginate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	responses := response.Paginate[wr.ToProjectByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
		Data: nil,
	}

	tx := app.Database.Where("`is_enable`=?", constant.IsEnableYes)

	if !strutil.IsEmpty(request.Classification) {
		tx = tx.Where("`classification_id`=?", request.Classification)
	}

	tx.Model(model.WebProject{}).Count(&responses.Total)

	if responses.Total > 0 {

		var projects []model.WebProject

		tx.Order("dated_at desc, `id` desc").Find(&projects)

		responses.Data = make([]wr.ToProjectByPaginate, len(projects))

		for index, item := range projects {
			responses.Data[index] = wr.ToProjectByPaginate{
				Id:      item.Id,
				Name:    item.Name,
				Picture: item.Picture,
				Address: item.Address,
				DatedAt: item.DatedAt.ToDateString(),
			}
		}
	}

	response.Success(ctx, responses)
}

func ToProjectByInformation(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.Success(ctx, errors.New("ID 不能为空"))
		return
	}

	var project model.WebProject

	fc := app.Database.Preload("Pictures").First(&project, "`id`=? and `is_enable`=?", id, constant.IsEnableYes)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "项目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("项目查找失败：%v", fc.Error))
		return
	}

	responses := wr.ToProjectByInformation{
		Id:             project.Theme,
		Theme:          project.Theme,
		Classification: project.ClassificationId,
		Name:           project.Name,
		Address:        project.Address,
		Picture:        project.Picture,
		Title:          project.Title,
		Keyword:        project.Keyword,
		Description:    project.Description,
		Html:           project.Html,
		DatedAt:        project.DatedAt.ToDateString(),
		Pictures:       make([]string, len(project.Pictures)),
	}

	for index, item := range project.Pictures {
		responses.Pictures[index] = item.Url
	}

	response.Success(ctx, responses)
}

func ToProjectByRelated(ctx *gin.Context) {

	var request web.ToProjectByRelated

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	tx := app.Database.Where("`is_enable`=?", constant.IsEnableYes)

	if !strutil.IsEmpty(request.Classification) {
		tx = tx.Where("`classification_id`=?", request.Classification)
	}

	var projects []model.WebProject

	tx.Order("rand()").Limit(8).Find(&projects)

	responses := make([]wr.ToProjectByPaginate, len(projects))

	for index, item := range projects {
		responses[index] = wr.ToProjectByPaginate{
			Id:      item.Id,
			Name:    item.Name,
			Picture: item.Picture,
			Address: item.Address,
			DatedAt: item.DatedAt.ToDateString(),
		}
	}

	response.Success(ctx, responses)
}
