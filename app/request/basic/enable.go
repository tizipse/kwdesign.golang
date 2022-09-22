package basic

type Enable struct {
	IsEnable int8 `form:"is_enable" json:"is_enable" binding:"required,oneof=1 2"`
}
