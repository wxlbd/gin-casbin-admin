package v1

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	menuApp "github.com/wxlbd/gin-casbin-admin/internal/application/menu"
	"github.com/wxlbd/gin-casbin-admin/internal/application/role"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type RoleHandler struct {
	svc role.Service
}

func NewRoleHandler(svc role.Service) *RoleHandler {
	return &RoleHandler{svc: svc}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req role.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	if err := h.svc.Create(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *RoleHandler) Update(c *gin.Context) {
	var req role.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
		return
	}
	req.ID = id
	if err := h.svc.Update(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	var ids []uint64
	for _, idStr := range idsStr {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
			return
		}
		ids = append(ids, id)
	}
	if err := h.svc.Delete(c, ids...); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *RoleHandler) Detail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
		return
	}
	resp, err := h.svc.FindByID(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, resp)
}

func (h *RoleHandler) List(c *gin.Context) {
	var req role.RoleListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	resp, err := h.svc.List(c, &req)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, resp)
}

func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	resp, err := h.svc.GetAllRoles(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, resp)
}

func (h *RoleHandler) AssignRoleMenusByIDs(c *gin.Context) {
	var req role.AssignRoleMenuIdsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
		return
	}
	if err := h.svc.AssignMenuByIds(c, id, req.MenuIds); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *RoleHandler) GetPermittedMenus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的角色ID"))
		return
	}
	menus, err := h.svc.GetPermittedMenus(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, menuApp.ToSysMenuList(menus))
}
