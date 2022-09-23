package web

import "kwd/app/request/basic"

type DoCategoryByCreate struct {
	Theme       string `json:"theme" form:"theme" binding:"required,oneof=light dark" label:"主题"`
	Uri         string `json:"uri" form:"uri" binding:"required,alpha,max=32" label:"URI"`
	Name        string `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Picture     string `json:"picture" form:"picture" binding:"omitempty,url,max=255" label:"图片"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Html        string `json:"html" form:"html" binding:"omitempty" label:"内容"`
	basic.Enable
}

type DoCategoryByUpdate struct {
	Theme       string `json:"theme" form:"theme" binding:"required,oneof=light dark" label:"主题"`
	Name        string `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Picture     string `json:"picture" form:"picture" binding:"omitempty,url,max=255" label:"图片"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Html        string `json:"html" form:"html" binding:"omitempty" label:"内容"`
	basic.Enable
}

type DoCategoryByEnable struct {
	Id int `form:"id" json:"id" binding:"required,gt=0"`
	basic.Enable
}
