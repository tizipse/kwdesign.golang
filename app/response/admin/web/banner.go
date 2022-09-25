package web

type ToBannerByPaginate struct {
	Id        int    `json:"id"`
	Theme     string `json:"theme"`
	Picture   string `json:"picture"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Url       string `json:"url"`
	IsEnable  int8   `json:"is_enable"`
	Order     int8   `json:"order"`
	CreatedAt string `json:"created_at"`
}
