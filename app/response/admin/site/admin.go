package site

type ToAdminByPaginate struct {
	Id        int                        `json:"id"`
	Username  string                     `json:"username"`
	Nickname  string                     `json:"nickname"`
	Mobile    string                     `json:"mobile"`
	Roles     []ToAdminByPaginateOfRoles `json:"roles"`
	IsEnable  int8                       `json:"is_enable"`
	CreatedAt string                     `json:"created_at"`
}

type ToAdminByPaginateOfRoles struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
