package v1

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/application/menu"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type MenuHandler struct {
	svc menu.Service
}

func NewMenuHandler(svc menu.Service) *MenuHandler {
	return &MenuHandler{svc: svc}
}

func (h *MenuHandler) Create(c *gin.Context) {
	var req menu.SysMenuRequest
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

func (h *MenuHandler) Update(c *gin.Context) {
	var req menu.SysMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的菜单ID"))
		return
	}
	req.ID = id
	if err := h.svc.Update(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *MenuHandler) Delete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	var ids []int64
	for _, idStr := range idsStr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的菜单ID"))
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

func (h *MenuHandler) List(c *gin.Context) {
	// Note: List request params are not fully implemented in service List method yet (it returns all).
	// But we can bind them anyway.
	var req menu.SysMenuListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	resp, err := h.svc.List(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, ginx.ListData{
		List:  resp,
		Total: int64(len(resp)), // Approximate total since we are fetching all
	})
}

func (h *MenuHandler) GetMenuTree(c *gin.Context) {
	tree, err := h.svc.GetMenuTree(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, tree)
}

func (h *MenuHandler) GetUserMenuTree(c *gin.Context) {
	userID := c.GetUint64("user_id")
	tree, err := h.svc.GetUserMenuTree(c, userID)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, tree)
}
