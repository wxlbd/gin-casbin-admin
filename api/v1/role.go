package v1

import "time"

type RoleAddRequest struct {
	RoleName      string   `json:"roleName"`
	RoleTag       string   `json:"roleTag"`
	Status        int      `json:"status"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permissionIds"`
}

type RoleUpdateRequest struct {
	Id            int      `json:"id"`
	RoleName      string   `json:"roleName"`
	RoleTag       string   `json:"roleTag"`
	Status        int      `json:"status"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permissionIds"`
}

type RoleDeleteRequest struct {
	Id int `json:"id"`
}

type RoleGetRequest struct {
	Id int `json:"id"`
}

type RoleListRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type RoleGetResponse struct {
	Id          int               `json:"id"`
	RoleName    string            `json:"role_name"`
	Status      int               `json:"status"`
	Description string            `json:"description"`
	CreatedAt   time.Time         `json:"created_at,omitempty"`
	Permissions []*RolePermission `json:"permissions,omitempty"`
}

type RoleListResponse struct {
	Total int
	List  []*RoleGetResponse
}

type RolePermission struct {
	Id     int    `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Name   string `json:"name"`
	Title  string `json:"title"`
}
