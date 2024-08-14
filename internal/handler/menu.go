package handler

import (
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type MenuHandler struct {
	*Handler
	menuService service.MenuService
}

func NewMenuHandler(handler *Handler, menuService service.MenuService) *MenuHandler {
	return &MenuHandler{
		Handler:     handler,
		menuService: menuService,
	}
}

func (h *MenuHandler) GetMenu(ctx *gin.Context) {

	req := new(v1.MenuGetRequest)
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	req.Id = uint64(id)
	resp, err := h.menuService.Get(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, resp)
}

func (h *MenuHandler) GetMenuList(ctx *gin.Context) {
	req := new(v1.MenuListRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	pageNumStr := ctx.DefaultQuery("pageNum", "1")
	req.PageNum, _ = strconv.Atoi(pageNumStr)
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	req.PageSize, _ = strconv.Atoi(pageSizeStr)
	resp, err := h.menuService.List(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, resp)
}

func (h *MenuHandler) GetMenuTree(ctx *gin.Context) {
	req := new(v1.MenuTreeRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	resp, err := h.menuService.Tree(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, resp)
}

func (h *MenuHandler) AddMenu(ctx *gin.Context) {
	req := new(v1.MenuAddRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err := h.menuService.Add(ctx, req)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *MenuHandler) UpdateMenu(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	req := new(v1.MenuUpdateRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	req.Id = uint64(id)
	if err := h.menuService.Update(ctx, req); err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}

func (h *MenuHandler) DeleteMenu(ctx *gin.Context) {
	req := new(v1.MenuDeleteRequest)
	if err := ctx.ShouldBindQuery(req); err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	ids := strings.Split(req.Ids, ",")
	if len(ids) == 0 {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}
	var intIds []uint64
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
			return
		}
		intIds = append(intIds, uint64(idInt))
	}
	err := h.menuService.Delete(ctx, intIds)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}
	v1.HandleSuccess(ctx, nil)
}
