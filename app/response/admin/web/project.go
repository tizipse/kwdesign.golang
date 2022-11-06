package web

type ToProjectByPaginate struct {
	Id             string `json:"id"`
	Classification string `json:"classification"`
	Theme          string `json:"theme"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Picture        string `json:"picture"`
	DatedAt        string `json:"dated_at,omitempty"`
	IsEnable       int8   `json:"is_enable"`
	CreatedAt      string `json:"created_at"`
}

type ToProjectByInformation struct {
	Id             string   `json:"id"`
	Classification string   `json:"classification"`
	Theme          string   `json:"theme"`
	Name           string   `json:"name"`
	Address        string   `json:"address"`
	Height         int8     `json:"height"`
	Picture        string   `json:"picture"`
	Title          string   `json:"title"`
	Keyword        string   `json:"keyword"`
	Description    string   `json:"description"`
	DatedAt        string   `json:"dated_at,omitempty"`
	Html           string   `json:"html"`
	Pictures       []string `json:"pictures"`
	IsEnable       int8     `json:"is_enable"`
}
