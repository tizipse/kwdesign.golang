package web

type ToCategories struct {
	Id                int    `json:"id"`
	Uri               string `json:"uri"`
	Name              string `json:"name"`
	IsRequiredPicture int8   `json:"is_required_picture"`
	Picture           string `json:"picture"`
	IsRequiredHtml    int8   `json:"is_required_html"`
	IsEnable          int8   `json:"is_enable"`
	CreatedAt         string `json:"created_at"`
}

type ToCategoryByInformation struct {
	Id                int    `json:"id"`
	Uri               string `json:"uri"`
	Name              string `json:"name"`
	Title             string `json:"title"`
	Keyword           string `json:"keyword"`
	Description       string `json:"description"`
	IsRequiredPicture int8   `json:"is_required_picture"`
	Picture           string `json:"picture"`
	IsRequiredHtml    int8   `json:"is_required_html"`
	Html              string `json:"html"`
	IsEnable          int8   `json:"is_enable"`
}
