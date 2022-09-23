package web

type ToContactByPaginate struct {
	Id        int    `json:"id"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Order     int8   `json:"order"`
	IsEnable  int8   `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
