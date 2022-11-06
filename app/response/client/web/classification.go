package web

type ToClassifications struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
}
