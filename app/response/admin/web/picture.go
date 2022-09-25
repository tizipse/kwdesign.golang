package web

type ToPictures struct {
	Id        int    `json:"id"`
	Label     string `json:"label"`
	Key       string `json:"key"`
	Val       string `json:"val"`
	Required  int8   `json:"required"`
	CreatedAt string `json:"created_at"`
}
