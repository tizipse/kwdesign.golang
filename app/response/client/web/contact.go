package web

type ToContacts struct {
	Id        int    `json:"id"`
	City      string `json:"city"`
	Address   string `json:"address"`
	Telephone string `json:"telephone,omitempty"`
}
