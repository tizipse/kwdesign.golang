package basic

import (
	"github.com/gin-gonic/gin"
	"kwd/app/request/admin/basic"
	res "kwd/app/response/admin/basic"
	"kwd/kernel/response"
	"kwd/kit/filesystem"
)

func DoUploadBySimple(ctx *gin.Context) {

	var request basic.DoUploadBySimple

	if err := ctx.ShouldBind(&request); err != nil {
		response.FailByRequest(ctx, err)
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		response.Fail(ctx, "上传失败，请稍后重试")
		return
	}

	storage := filesystem.New().Upload()

	uri, name, err := storage.Save(file, request.Dir, "")

	if err != nil {
		response.Fail(ctx, "上传失败，请稍后重试")
		return
	}

	responses := res.DoUploadBySimple{
		Name: name,
		Path: uri,
		Url:  storage.Url(uri),
	}

	response.Success(ctx, responses)
}
