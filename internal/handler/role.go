package handler

import (
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	*Handler
	roleService service.RoleService
}

func NewRoleHandler(handler *Handler, roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		Handler:     handler,
		roleService: roleService,
	}
}
func (r *RoleHandler) Add(ctx *gin.Context) {
	req := new(v1.RoleAddRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	if err := r.roleService.Add(ctx, req); err != nil {
		r.logger.WithContext(ctx).Error("roleService.Add error", zap.Error(err))
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (r *RoleHandler) Delete(ctx *gin.Context) {
	req := new(v1.RoleDeleteRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	if err := r.roleService.Delete(ctx, req); err != nil {
		r.logger.WithContext(ctx).Error("roleService.Delete error", zap.Error(err))
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (r *RoleHandler) Update(ctx *gin.Context) {
	req := new(v1.RoleUpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	if err := r.roleService.Update(ctx, req); err != nil {
		r.logger.WithContext(ctx).Error("roleService.Update error", zap.Error(err))
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (r *RoleHandler) Get(ctx *gin.Context) {
	id, err := GetPathParamInt(ctx, "id")
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	resp, err := r.roleService.Get(ctx, &v1.RoleGetRequest{Id: id})
	if err != nil {
		r.logger.WithContext(ctx).Error("roleService.Get error", zap.Error(err))
		return
	}
	v1.HandleSuccess(ctx, resp)
}

func (r *RoleHandler) List(ctx *gin.Context) {
	req := new(v1.RoleListRequest)
	req.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	req.Page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	list, err := r.roleService.List(ctx, req)
	if err != nil {
		r.logger.WithContext(ctx).Error("roleService.List error", zap.Error(err))
		return
	}
	v1.HandleSuccess(ctx, list)
}
