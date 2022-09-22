package api

import (
	"net/http"
)

var (
	cache       bool
	OmitOfCache = make(map[string]bool, 0)
)

func InitApi() {
	if !cache && len(Omits()) > 0 {
		for _, item := range Omits() {
			OmitOfCache[OmitKey(item.Method, item.Path)] = true
		}
		cache = true
	}
}

func OmitKey(method string, path string) string {
	return method + ":" + path
}

// Omits 被忽略返回的 Api
func Omits() []Api {
	return []Api{

		{Method: http.MethodPost, Path: "/admin/upload"},
		{Method: http.MethodPost, Path: "/admin/login/account"},
		{Method: http.MethodPost, Path: "/admin/login/qrcode"},
		{Method: http.MethodGet, Path: "/admin/account/information"},
		{Method: http.MethodGet, Path: "/admin/account/module"},
		{Method: http.MethodGet, Path: "/admin/account/permission"},
		{Method: http.MethodPost, Path: "/admin/account/logout"},

		{Method: http.MethodGet, Path: "/admin/site/helper/apis"},
		{Method: http.MethodGet, Path: "/admin/site/permission/parents"},
		{Method: http.MethodGet, Path: "/admin/site/permission/self"},
		{Method: http.MethodGet, Path: "/admin/site/role/enable"},
		{Method: http.MethodGet, Path: "/admin/site/module/enable"},

		{Method: http.MethodGet, Path: "/admin/kwd/category/parents"},
		{Method: http.MethodGet, Path: "/admin/kwd/categories/:id"},
	}
}
