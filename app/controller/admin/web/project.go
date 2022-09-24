package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/model"
	"kwd/app/request/admin/web"
	wr "kwd/app/response/admin/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func DoProjectByCreate(ctx *gin.Context) {

	var request web.DoProjectByCreate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var classification model.WebClassification

	fcl := app.Database.First(&classification, request.Classification)

	if errors.Is(fcl.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "分类不存在")
		return
	} else if fcl.Error != nil {
		response.Fail(ctx, fmt.Sprintf("分类查找失败：%v", fcl.Error))
		return
	}

	now := carbon.Now()

	project := model.WebProject{
		Id:               app.Snowflake.Generate().String(),
		ClassificationId: request.Classification,
		Theme:            request.Theme,
		Name:             request.Name,
		Address:          request.Address,
		Picture:          request.Picture,
		Title:            request.Title,
		Keyword:          request.Keyword,
		Description:      request.Description,
		DatedAt:          carbon.Date{Carbon: now},
		Html:             request.Html,
		IsEnable:         request.IsEnable,
	}

	if !strutil.IsEmpty(request.Date) {

		date := carbon.ParseByFormat(request.Date, "Y-m-d")

		if !date.IsZero() {
			project.DatedAt = carbon.Date{Carbon: date}
		}
	}

	tx := app.Database.Begin()

	if cp := tx.Create(&project); cp.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("项目创建失败：%v", cp.Error))
		return
	}

	if len(request.Pictures) > 0 {

		pictures := make([]model.WebProjectPicture, len(request.Pictures))

		for index, item := range request.Pictures {
			pictures[index] = model.WebProjectPicture{
				ClassificationId: project.ClassificationId,
				ProjectId:        project.Id,
				Url:              item,
			}
		}

		if cp := tx.Create(&pictures); cp.Error != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("项目创建失败：%v", cp.Error))
			return
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoProjectByUpdate(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var request web.DoProjectByUpdate

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var project model.WebProject

	fc := app.Database.Preload("Pictures").First(&project, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "项目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("项目查找失败：%v", fc.Error))
		return
	}

	if request.Classification != project.ClassificationId {

		var classification model.WebClassification

		fcl := app.Database.First(&classification, request.Classification)

		if errors.Is(fcl.Error, gorm.ErrRecordNotFound) {
			response.NotFound(ctx, "分类不存在")
			return
		} else if fcl.Error != nil {
			response.Fail(ctx, fmt.Sprintf("分类查找失败：%v", fcl.Error))
			return
		}
	}

	creates := make([]model.WebProjectPicture, 0)
	deletes := make([]int, 0)

	for _, item := range request.Pictures {

		mark := true

		for _, value := range project.Pictures {
			if item == value.Url {
				mark = false
			}
		}

		if mark {
			creates = append(creates, model.WebProjectPicture{
				ClassificationId: project.ClassificationId,
				ProjectId:        project.Id,
				Url:              item,
			})
		}
	}

	for _, item := range project.Pictures {

		mark := true

		for _, value := range request.Pictures {
			if item.Url == value {
				mark = false
			}
		}

		if mark {
			deletes = append(deletes, item.Id)
		}
	}

	project.ClassificationId = request.Classification
	project.Theme = request.Theme
	project.Name = request.Name
	project.Address = request.Address
	project.Picture = request.Picture
	project.Title = request.Title
	project.Keyword = request.Keyword
	project.Description = request.Description
	project.Html = request.Html
	project.IsEnable = request.IsEnable

	if !strutil.IsEmpty(request.Date) {

		date := carbon.ParseByFormat(request.Date, "Y-m-d")

		if !date.IsZero() {
			project.DatedAt = carbon.Date{Carbon: date}
		}
	}

	tx := app.Database.Begin()

	if up := tx.Save(&project); up.Error != nil {
		tx.Rollback()
		response.Fail(ctx, fmt.Sprintf("项目修改失败：%v", up.Error))
		return
	}

	if len(creates) > 0 {

		if cp := tx.Create(&creates); cp.Error != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("项目修改失败：%v", cp.Error))
			return
		}
	}

	if len(deletes) > 0 {

		if cp := tx.Delete(&model.WebProjectPicture{}, deletes); cp.Error != nil {
			tx.Rollback()
			response.Fail(ctx, fmt.Sprintf("项目修改失败：%v", cp.Error))
			return
		}
	}

	tx.Commit()

	response.Success[any](ctx)
}

func DoProjectByDelete(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var project model.WebProject

	fp := app.Database.First(&project, id)

	if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "项目不存在")
		return
	} else if fp.Error != nil {
		response.Fail(ctx, fmt.Sprintf("项目查找失败：%v", fp.Error))
		return
	}

	if dc := app.Database.Delete(&project); dc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("项目删除失败：%v", dc.Error))
		return
	}

	response.Success[any](ctx)
}

func DoProjectByEnable(ctx *gin.Context) {

	var request web.DoProjectByEnable

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	var category model.WebProject

	fc := app.Database.First(&category, request.Id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "项目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("项目查找失败：%v", fc.Error))
		return
	}

	category.IsEnable = request.IsEnable

	if uc := app.Database.Save(&category); uc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("启禁失败：%v", uc.Error))
		return
	}

	response.Success[any](ctx)
}

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

	tx := app.Database.Model(model.WebProject{})

	tx.Count(&responses.Total)

	if responses.Total > 0 {

		var projects []model.WebProject

		tx.
			Preload("Classification", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`id` desc").
			Find(&projects)

		responses.Data = make([]wr.ToProjectByPaginate, len(projects))

		for index, item := range projects {
			responses.Data[index] = wr.ToProjectByPaginate{
				Id:        item.Id,
				Theme:     item.Theme,
				Name:      item.Name,
				Address:   item.Address,
				Picture:   item.Picture,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
			if item.Classification != nil {
				responses.Data[index].Classification = item.Classification.Name
			}
		}
	}

	response.Success(ctx, responses)
}

func ToProjectByInformation(ctx *gin.Context) {

	id := ctx.Param("id")

	if strutil.IsEmpty(id) {
		response.FailByRequest(ctx, errors.New("ID 不存在"))
		return
	}

	var project model.WebProject

	fc := app.Database.
		Preload("Pictures").
		First(&project, id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "项目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("项目查找失败：%v", fc.Error))
		return
	}

	responses := wr.ToProjectByInformation{
		Id:             project.Id,
		Classification: project.ClassificationId,
		Theme:          project.Theme,
		Name:           project.Name,
		Address:        project.Address,
		Picture:        project.Picture,
		Title:          project.Title,
		Keyword:        project.Keyword,
		Description:    project.Description,
		Html:           project.Html,
		DatedAt:        project.DatedAt.ToDateString(),
		Pictures:       make([]string, len(project.Pictures)),
		IsEnable:       project.IsEnable,
	}

	for index, item := range project.Pictures {
		responses.Pictures[index] = item.Url
	}

	response.Success(ctx, responses)

}
