package web

type ToBanners struct {
	Id      int    `json:"id"`
	Client  string `json:"client"`
	Theme   string `json:"theme"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
	Target  string `json:"target"`
	Url     string `json:"url"`
}
