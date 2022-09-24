package web

import "kwd/app/request/basic"

type DoClassificationByCreate struct {
	Name        string `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Order       int8   `form:"order" json:"order" binding:"omitempty,gt=1,lt=99" label:"序号"`
	basic.Enable
}

type DoClassificationByUpdate struct {
	Name        string `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Order       int8   `form:"order" json:"order" binding:"omitempty,gt=1,lt=99" label:"序号"`
	basic.Enable
}

type DoClassificationByEnable struct {
	Id string `form:"id" json:"id" binding:"required,snowflake" label:"ID"`
	basic.Enable
}
