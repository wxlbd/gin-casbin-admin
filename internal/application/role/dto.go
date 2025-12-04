package role

import (
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// RoleRequest 创建/更新角色请求
type RoleRequest struct {
	ID     uint64 `json:"id"` // 更新时必填
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Status any    `json:"status"`
	Sort   int16  `json:"sort"`
	Remark string `json:"remark"`
}

func (req *RoleRequest) ToEntity() *role.Role {
	status := int8(1) // Default to 1 (Normal)

	switch v := req.Status.(type) {
	case string:
		if v == "0" {
			status = 0
		}
	case float64:
		status = int8(v)
	case int:
		status = int8(v)
	}

	return &role.Role{
		ID:     req.ID,
		Name:   req.Name,
		Code:   req.Code,
		Status: status,
		Sort:   req.Sort,
		Remark: req.Remark,
	}
}

// RoleListRequest 角色列表请求
type RoleListRequest struct {
	*types.PageParam
	Name   string `form:"name"`
	Code   string `form:"code"`
	Status int8   `form:"status"`
}

func (req *RoleListRequest) ToQuery() *role.RoleQuery {
	return &role.RoleQuery{
		PageParam: req.PageParam,
		Name:      req.Name,
		Code:      req.Code,
		Status:    req.Status,
	}
}

// RoleResponse 角色信息响应
type RoleResponse struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Status  int8   `json:"status"`
	Sort    int16  `json:"sort"`
	Remark  string `json:"remark"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func ToRoleResponse(r *role.Role) *RoleResponse {
	if r == nil {
		return nil
	}
	return &RoleResponse{
		ID:      r.ID,
		Name:    r.Name,
		Code:    r.Code,
		Status:  r.Status,
		Sort:    r.Sort,
		Remark:  r.Remark,
		Created: r.CreatedAt.Format(time.DateTime),
		Updated: r.UpdatedAt.Format(time.DateTime),
	}
}

func ToRoleList(roles []*role.Role) []*RoleResponse {
	list := make([]*RoleResponse, 0, len(roles))
	for _, r := range roles {
		list = append(list, ToRoleResponse(r))
	}
	return list
}

// RoleListResponse 角色列表响应
type RoleListResponse struct {
	List  []*RoleResponse `json:"list"`
	Total int64           `json:"total"`
}

// AssignRoleMenuIdsRequest 分配菜单请求
type AssignRoleMenuIdsRequest struct {
	MenuIds []uint64 `json:"menuIds"`
}
