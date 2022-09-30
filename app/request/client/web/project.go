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
