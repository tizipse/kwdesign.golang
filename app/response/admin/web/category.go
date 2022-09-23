package web

type ToCategories struct {
	Id        int    `json:"id"`
	Theme     string `json:"theme"`
	Uri       string `json:"uri"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsEnable  int8   `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToCategoryByInformation struct {
	Id          int    `json:"id"`
	Theme       string `json:"theme"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Html        string `json:"html"`
	IsEnable    int8   `json:"is_enable"`
}
