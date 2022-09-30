package web

type ToProjectByPaginate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Picture string `json:"picture"`
	DatedAt string `json:"dated_at"`
}

type ToProjectByInformation struct {
	Id             string   `json:"id"`
	Theme          string   `json:"theme"`
	Classification string   `json:"classification"`
	Name           string   `json:"name"`
	Address        string   `json:"address"`
	Picture        string   `json:"picture"`
	Title          string   `json:"title"`
	Keyword        string   `json:"keyword"`
	Description    string   `json:"description"`
	Html           string   `json:"html"`
	DatedAt        string   `json:"dated_at"`
	Pictures       []string `json:"pictures"`
}

type ToProjectByRelated struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Picture string `json:"picture"`
	DatedAt string `json:"dated_at"`
}
