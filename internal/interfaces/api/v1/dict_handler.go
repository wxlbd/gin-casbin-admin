package v1

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/application/dict"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type DictHandler struct {
	svc dict.Service
}

func NewDictHandler(svc dict.Service) *DictHandler {
	return &DictHandler{svc: svc}
}

// --- DictType ---

func (h *DictHandler) CreateDictType(c *gin.Context) {
	var req dict.DictTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	if err := h.svc.CreateType(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *DictHandler) UpdateDictType(c *gin.Context) {
	var req dict.DictTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}
	req.ID = id
	if err := h.svc.UpdateType(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *DictHandler) DeleteDictType(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	var ids []int64
	for _, idStr := range idsStr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
			return
		}
		ids = append(ids, id)
	}
	if err := h.svc.DeleteType(c, ids...); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *DictHandler) GetDictType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}
	resp, err := h.svc.GetType(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, resp)
}

func (h *DictHandler) ListDictType(c *gin.Context) {
	var req dict.DictTypeListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	resp, err := h.svc.ListType(c, &req)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, ginx.ListData{
		List:  resp.List,
		Total: resp.Total,
	})
}

// --- DictData ---

func (h *DictHandler) CreateDictData(c *gin.Context) {
	var req dict.DictDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	if err := h.svc.CreateData(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *DictHandler) UpdateDictData(c *gin.Context) {
	var req dict.DictDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}
	req.ID = id
	if err := h.svc.UpdateData(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *DictHandler) DeleteDictData(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	var ids []int64
	for _, idStr := range idsStr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
			return
		}
		ids = append(ids, id)
	}
	if err := h.svc.DeleteData(c, ids...); err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, nil)
}

func (h *DictHandler) GetDictData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}
	resp, err := h.svc.GetData(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, resp)
}

func (h *DictHandler) ListDictData(c *gin.Context) {
	var req dict.DictDataListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}
	resp, err := h.svc.ListData(c, &req)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, ginx.ListData{
		List:  resp.List,
		Total: resp.Total,
	})
}

func (h *DictHandler) GetDictDataByType(c *gin.Context) {
	typeCode := c.Param("type")
	if typeCode == "" {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "字典类型编码不能为空"))
		return
	}
	// Note: Service doesn't have GetDictDataByType yet.
	// I should add it to service or use ListData with typeCode.
	// Let's use ListData.
	req := dict.DictDataListRequest{
		TypeCode: typeCode,
	}
	// We need to handle pagination or fetch all.
	// Assuming we want all for this endpoint.
	// But ListData might require pagination.
	// Let's just call ListData with default pagination or large size.
	// Or better, add FindAllByType to service.
	// For now, I'll use ListData with large page size.
	req.Page = 1
	req.PageSize = 1000

	resp, err := h.svc.ListData(c, &req)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}
	ginx.Success(c, resp.List)
}
