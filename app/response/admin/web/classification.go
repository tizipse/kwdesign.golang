package web

type ToClassifications struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Order     int8   `json:"order"`
	IsEnable  int8   `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToClassificationByInformation struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Order       int8   `json:"order"`
	IsEnable    int8   `json:"is_enable"`
}

type ToClassificationByEnable struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
