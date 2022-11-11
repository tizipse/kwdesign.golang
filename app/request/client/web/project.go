package web

import "kwd/app/request/basic"

type ToProjectByPaginate struct {
	Classification string `json:"classification" form:"classification" binding:"omitempty,snowflake"`

	basic.Paginate
}

type ToProjectByRelated struct {
	Classification string `json:"classification" form:"classification" binding:"required,snowflake"`
	Project        string `json:"project" form:"project" binding:"omitempty,snowflake"`
}

type ToProjectByRecommend struct {
	Number         int8     `json:"number" form:"number" binding:"required,gte=1,lte=20" label:"数量"`
	Classification string   `json:"classification" form:"classification" binding:"omitempty,snowflake" label:"分类"`
	Excludes       []string `json:"excludes" form:"excludes" binding:"omitempty,max=10,dive,snowflake" label:"排除"`
}
