package handler

import (
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type PermissionHandler struct {
	*Handler
	permissionService service.PermissionsService
}

func NewPermissionHandler(handler *Handler, permissionService service.PermissionsService) *PermissionHandler {
	return &PermissionHandler{
		Handler: handler,

		permissionService: permissionService,
	}
}

func (p *PermissionHandler) Add(ctx *gin.Context) {
	req := new(v1.PermissionsAddRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	if err := p.permissionService.Add(ctx, req); err != nil {
		p.logger.WithContext(ctx).Error("permissionService.Add error", zap.Error(err))
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (p *PermissionHandler) Delete(ctx *gin.Context) {
	req := new(v1.PermissionsDeleteRequest)
	var err error
	req.Id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	if err := p.permissionService.Delete(ctx, req); err != nil {
		p.logger.WithContext(ctx).Error("permissionService.Delete error", zap.Error(err))
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (p *PermissionHandler) Update(ctx *gin.Context) {
	req := new(v1.PermissionsUpdateRequest)
	var err error
	req.Id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	if err := p.permissionService.Update(ctx, req); err != nil {
		p.logger.WithContext(ctx).Error("permissionService.Update error", zap.Error(err))
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (p *PermissionHandler) Tree(ctx *gin.Context) {
	req := new(v1.PermissionTreeRequest)
	list, err := p.permissionService.List(ctx, req)
	if err != nil {
		p.logger.WithContext(ctx).Error("permissionService.List error", zap.Error(err))
		v1.HandleError(ctx, http.StatusOK, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, list)
}
