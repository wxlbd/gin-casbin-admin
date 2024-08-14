package v1

import "time"

type MenuAddRequest struct {
	Type          int8   `json:"type" binding:"oneof=1 2 3"`
	ParentId      uint64 `json:"parentId"`
	Icon          string `json:"icon"`
	Title         string `json:"title" binding:"required"`
	Path          string `json:"path"`                                                  // 路由地址
	Method        string `json:"method" binding:"oneof=GET POST PUT PATCH HEAD DELETE"` // 请求方式
	Redirect      string `json:"redirect"`                                              // 重定向地址
	ComponentName string `json:"componentName"`
	ComponentPath string `json:"componentPath"`
	IsHidden      int8   `json:"isHidden"`
	IsExternal    int8   `json:"isExternal"`
	Sort          int    `json:"sort"`
	Status        int    `json:"status"` // 状态 1 启用 2 禁用
}

type MenuUpdateRequest struct {
	Id            uint64 `json:"id"`
	Type          int8   `json:"type" binding:"oneof=1 2 3"`
	ParentId      uint64 `json:"parentId"`
	Icon          string `json:"icon"`
	Title         string `json:"title" binding:"required"`
	Path          string `json:"path"`                                                  // 路由地址
	Method        string `json:"method" binding:"oneof=GET POST PUT PATCH HEAD DELETE"` // 请求方式
	Redirect      string `json:"redirect"`                                              // 重定向地址
	ComponentName string `json:"componentName"`
	ComponentPath string `json:"componentPath"`
	IsHidden      int8   `json:"isHidden"`
	IsExternal    int8   `json:"isExternal"`
	Sort          int    `json:"sort"`
	Status        int    `json:"status"` // 状态 1 启用 2 禁用
}

type MenuDeleteRequest struct {
	Ids string `form:"ids" json:"ids"`
}

type MenuGetRequest struct {
	Id uint64 `form:"id" json:"id"`
}

type MenuListRequest struct {
	Title    string `json:"title"`
	Type     int    `json:"type"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type Menu struct {
	Id            uint64    `json:"id"`
	Type          int8      `json:"type"`
	ParentId      uint64    `json:"parentId"`
	Icon          string    `json:"icon"`
	Title         string    `json:"title"`
	Path          string    `json:"path"`     // 路由地址
	Method        string    `json:"method"`   // 请求方式
	Redirect      string    `json:"redirect"` // 重定向地址
	ComponentName string    `json:"componentName"`
	ComponentPath string    `json:"componentPath"` // 组件路径
	IsHidden      int8      `json:"isHidden"`
	IsExternal    int8      `json:"isExternal"`
	Sort          int       `json:"sort"`
	Status        int       `json:"status"` // 状态 1 启用 2 禁用
	Children      []*Menu   `json:"children,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type MenuListResponse struct {
	List  []*Menu `json:"list"`
	Total int64   `json:"total"`
}

type MenuGetResponse struct {
	*Menu
}

type MenuTreeRequest struct {
}
type MenuTreeResponse struct {
	List []*Menu `json:"list"`
}
