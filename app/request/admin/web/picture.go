package web

type DoPictureByCreate struct {
	Label    string `json:"label" form:"label" binding:"required,max=10"`
	Key      string `json:"key" form:"key" binding:"required,max=20,alpha|ascii"`
	Val      string `json:"val" form:"val" binding:"required,max=255,url"`
	Required int8   `json:"required" form:"required" binding:"required,oneof=1 2"`
}

type DoPictureByUpdate struct {
	Label string `json:"label" form:"label" binding:"required,max=10"`
	Val   string `json:"val" form:"val" binding:"omitempty,max=255,url"`
}
