package web

import "kwd/app/request/basic"

type DoBannerByCreate struct {
	Layout  string `json:"layout" form:"layout" binding:"required,oneof=mobile ipad pc" label:"布局"`
	Picture string `json:"picture" form:"picture" binding:"required,url,max=255" label:"图片"`
	Name    string `form:"name" json:"name" binding:"required,max=32" label:"名称"`
	Target  string `json:"target" form:"target" binding:"required,oneof=blank self" label:"打开"`
	Url     string `json:"url" form:"url" binding:"omitempty,url|uri,max=255" label:"链接"`
	Order   int8   `form:"order" json:"order" binding:"omitempty,gt=1,lt=99" label:"序号"`
	basic.Enable
}

type DoBannerByUpdate struct {
	Layout  string `json:"layout" form:"layout" binding:"required,oneof=mobile ipad pc" label:"布局"`
	Picture string `json:"picture" form:"picture" binding:"required,url,max=255" label:"图片"`
	Name    string `form:"name" json:"name" binding:"required,min=2,max=20" label:"名称"`
	Target  string `json:"target" form:"target" binding:"required,oneof=blank self" label:"打开"`
	Url     string `json:"url" form:"url" binding:"omitempty,url|uri,max=255" label:"链接"`
	Order   int8   `form:"order" json:"order" binding:"omitempty,gt=1,lt=99" label:"序号"`
	basic.Enable
}

type DoBannerByEnable struct {
	Id int `form:"id" json:"id" binding:"required,gt=0"`
	basic.Enable
}

type ToBannerByPaginate struct {
	basic.Paginate
}
