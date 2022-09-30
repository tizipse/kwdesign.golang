package web

type ToCategoryByInformation struct {
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Picture     string `json:"picture,omitempty"`
	Title       string `json:"title,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	Description string `json:"description,omitempty"`
	Html        string `json:"html,omitempty"`
}
