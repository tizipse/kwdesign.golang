package web

import (
	"kwd/app/request/basic"
)

type DoProjectByCreate struct {
	Classification string   `json:"classification" form:"classification" binding:"required,snowflake" label:"分类"`
	Theme          string   `json:"theme" form:"theme" binding:"required,oneof=light dark" label:"主题"`
	Name           string   `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Address        string   `json:"address" form:"address" binding:"omitempty,max=64" label:"地点"`
	Height         int8     `json:"height" form:"height" binding:"required,gte=1,lte=100" label:"高度"`
	Date           string   `json:"date" form:"date" binding:"omitempty,datetime=2006-01-02" label:"日期"`
	Picture        string   `json:"picture" form:"picture" binding:"omitempty,url,max=255" label:"图片"`
	Title          string   `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword        string   `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description    string   `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Html           string   `json:"html" form:"html" binding:"omitempty" label:"内容"`
	Pictures       []string `json:"pictures" form:"pictures" binding:"omitempty,max=16,dive,url,max=255" label:"图片"`
	basic.Enable
}

type DoProjectByUpdate struct {
	Classification string   `json:"classification" form:"classification" binding:"required,snowflake" label:"分类"`
	Theme          string   `json:"theme" form:"theme" binding:"required,oneof=light dark" label:"主题"`
	Name           string   `json:"name" form:"name" binding:"required,max=32" label:"名称"`
	Address        string   `json:"address" form:"address" binding:"omitempty,max=64" label:"地点"`
	Height         int8     `json:"height" form:"height" binding:"required,gte=1,lte=100" label:"高度"`
	Date           string   `json:"date" form:"date" binding:"omitempty,datetime=2006-01-02" label:"日期"`
	Picture        string   `json:"picture" form:"picture" binding:"omitempty,url,max=255" label:"图片"`
	Title          string   `json:"title" form:"title" binding:"omitempty,max=255" label:"SEO 标题"`
	Keyword        string   `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"SEO 关键词"`
	Description    string   `json:"description" form:"description" binding:"omitempty,max=255" label:"SEO 描述"`
	Html           string   `json:"html" form:"html" binding:"omitempty" label:"内容"`
	Pictures       []string `json:"pictures" form:"pictures" binding:"omitempty,max=16,dive,url,max=255" label:"图片"`
	basic.Enable
}

type DoProjectByEnable struct {
	Id string `form:"id" json:"id" binding:"required,snowflake"`
	basic.Enable
}

type ToProjectByPaginate struct {
	basic.Paginate
}
