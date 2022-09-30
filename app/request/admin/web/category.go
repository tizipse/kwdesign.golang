package web

import "kwd/app/request/basic"

type DoCategoryByCreate struct {
	Uri               string `json:"uri" form:"uri" binding:"required,alpha,max=32" label:"URI"`
	Name              string `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Title             string `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword           string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description       string `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	IsRequiredPicture int8   `json:"is_required_picture" form:"is_required_picture" binding:"required,oneof=1 2" label:"是否必填图片"`
	IsRequiredHtml    int8   `json:"is_required_html" form:"is_required_html" binding:"required,oneof=1 2" label:"是否必填内容"`
	Picture           string `json:"picture" form:"picture" binding:"required_if=IsRequiredPicture 1,url,max=255" label:"图片"`
	Html              string `json:"html" form:"html" binding:"required_if=IsRequiredHtml 1" label:"内容"`
	basic.Enable
}

type DoCategoryByUpdate struct {
	Name        string `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Picture     string `json:"picture" form:"picture" binding:"omitempty,url,max=255" label:"图片"`
	Html        string `json:"html" form:"html" binding:"omitempty" label:"内容"`
	basic.Enable
}

type DoCategoryByEnable struct {
	Id int `form:"id" json:"id" binding:"required,gt=0"`
	basic.Enable
}

type DoCategoryByIsRequiredPicture struct {
	Id                int  `form:"id" json:"id" binding:"required,gt=0"`
	IsRequiredPicture int8 `json:"is_required_picture" form:"is_required_picture" binding:"required,oneof=1 2" label:"是否必填图片"`
}

type DoCategoryByIsRequiredHtml struct {
	Id             int  `form:"id" json:"id" binding:"required,gt=0"`
	IsRequiredHtml int8 `json:"is_required_html" form:"is_required_html" binding:"required,oneof=1 2" label:"是否必填内容"`
}
