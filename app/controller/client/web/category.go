package web

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
	"kwd/app/constant"
	"kwd/app/model"
	wr "kwd/app/response/client/web"
	"kwd/kernel/app"
	"kwd/kernel/response"
)

func ToCategoryByInformation(ctx *gin.Context) {

	uri := ctx.Param("uri")

	if strutil.IsEmpty(uri) {
		response.FailByRequest(ctx, errors.New("URI 不能为空"))
		return
	}

	var category model.WebCategory

	fc := app.Database.First(&category, "`uri`=? and `is_enable`=?", uri, constant.IsEnableYes)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		response.NotFound(ctx, "栏目不存在")
		return
	} else if fc.Error != nil {
		response.Fail(ctx, fmt.Sprintf("栏目查找失败：%v", fc.Error))
		return
	}

	responses := wr.ToCategoryByInformation{
		Uri:         category.Uri,
		Name:        category.Name,
		Picture:     category.Picture,
		Title:       category.Title,
		Keyword:     category.Keyword,
		Description: category.Description,
		Html:        category.Html,
	}

	response.Success(ctx, responses)
}
