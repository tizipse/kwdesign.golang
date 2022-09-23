package web

import "kwd/app/request/basic"

type DoContactByCreate struct {
	City      string `json:"city" form:"city" binding:"required,max=32"`
	Address   string `json:"address" form:"address" binding:"required,max=255"`
	Telephone string `json:"telephone" form:"telephone" binding:"required,max=32"`
	Order     int8   `json:"order" form:"order" binding:"required,gte=1,lte=99"`
	IsEnable  int8   `json:"is_enable" form:"is_enable" binding:"required,oneof=1 2"`
}

type DoContactByUpdate struct {
	City      string `json:"city" form:"city" binding:"required,max=32"`
	Address   string `json:"address" form:"address" binding:"required,max=255"`
	Telephone string `json:"telephone" form:"telephone" binding:"required,max=32"`
	Order     int8   `json:"order" form:"order" binding:"required,gte=1,lte=99"`
	IsEnable  int8   `json:"is_enable" form:"is_enable" binding:"required,oneof=1 2"`
}

type DoContactByEnable struct {
	Id int `form:"id" json:"id" binding:"required,gt=0"`
	basic.Enable
}

type ToContactByPaginate struct {
	basic.Paginate
}
