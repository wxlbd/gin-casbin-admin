package v1

import "time"

type RoleAddRequest struct {
	RoleName    string   `json:"roleName"`
	RoleTag     string   `json:"roleTag"`
	Status      int      `json:"status"`
	Description string   `json:"description"`
	MenuIds     []string `json:"menuIds"`
}

type RoleUpdateRequest struct {
	Id          int      `json:"id"`
	RoleName    string   `json:"roleName"`
	RoleTag     string   `json:"roleTag"`
	Status      int      `json:"status"`
	Description string   `json:"description"`
	MenuIds     []string `json:"menuIds"`
}

type RoleDeleteRequest struct {
	Id int `json:"id"`
}

type RoleGetRequest struct {
	Id int `json:"id"`
}

type RoleListRequest struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

type RoleGetResponse struct {
	Id          int         `json:"id"`
	RoleName    string      `json:"roleName"`
	Status      int         `json:"status"`
	Description string      `json:"description"`
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
	Menus       []*RoleMenu `json:"menus,omitempty"`
}

type RoleListResponse struct {
	Total int
	List  []*RoleGetResponse
}

type RoleMenu struct {
	Id     int    `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Title  string `json:"title"`
}
