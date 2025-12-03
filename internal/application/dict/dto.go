package dict

import (
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/dict"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// DictType DTOs
type DictTypeRequest struct {
	ID     int64  `json:"id"`
	Name   string `json:"name" binding:"required"`
	Code   string `json:"code" binding:"required"`
	Status int32  `json:"status" binding:"required"`
	Sort   int32  `json:"sort"`
	Remark string `json:"remark"`
}

func (req *DictTypeRequest) ToEntity() *dict.DictType {
	return &dict.DictType{
		ID:     req.ID,
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
		Sort:   req.Sort,
		Remark: req.Remark,
	}
}

type DictTypeResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Status    int32     `json:"status"`
	Sort      int32     `json:"sort"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDictTypeResponse(d *dict.DictType) *DictTypeResponse {
	if d == nil {
		return nil
	}
	return &DictTypeResponse{
		ID:        d.ID,
		Name:      d.Name,
		Code:      d.Code,
		Status:    d.Status,
		Sort:      d.Sort,
		Remark:    d.Remark,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

type DictTypeListRequest struct {
	*types.PageParam
	Name   string `form:"name"`
	Code   string `form:"code"`
	Status int32  `form:"status"`
}

func (req *DictTypeListRequest) ToQuery() *dict.DictTypeQuery {
	return &dict.DictTypeQuery{
		PageParam: req.PageParam,
		Name:      req.Name,
		Code:      req.Code,
		Status:    req.Status,
	}
}

type DictTypeListResponse struct {
	List  []*DictTypeResponse `json:"list"`
	Total int64               `json:"total"`
}

// DictData DTOs
type DictDataRequest struct {
	ID       int64  `json:"id"`
	TypeCode string `json:"typeCode" binding:"required"`
	Label    string `json:"label" binding:"required"`
	Value    string `json:"value" binding:"required"`
	Status   int32  `json:"status" binding:"required"`
	Sort     int32  `json:"sort"`
	Remark   string `json:"remark"`
}

func (req *DictDataRequest) ToEntity() *dict.DictData {
	return &dict.DictData{
		ID:       req.ID,
		TypeCode: req.TypeCode,
		Label:    req.Label,
		Value:    req.Value,
		Status:   req.Status,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}
}

type DictDataResponse struct {
	ID        int64     `json:"id"`
	TypeCode  string    `json:"type_code"`
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	Status    int32     `json:"status"`
	Sort      int32     `json:"sort"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDictDataResponse(d *dict.DictData) *DictDataResponse {
	if d == nil {
		return nil
	}
	return &DictDataResponse{
		ID:        d.ID,
		TypeCode:  d.TypeCode,
		Label:     d.Label,
		Value:     d.Value,
		Status:    d.Status,
		Sort:      d.Sort,
		Remark:    d.Remark,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

type DictDataListRequest struct {
	*types.PageParam
	TypeCode string `form:"typeCode"`
	Label    string `form:"label"`
	Status   int32  `form:"status"`
}

func (req *DictDataListRequest) ToQuery() *dict.DictDataQuery {
	return &dict.DictDataQuery{
		PageParam: req.PageParam,
		TypeCode:  req.TypeCode,
		Label:     req.Label,
		Status:    req.Status,
	}
}

type DictDataListResponse struct {
	List  []*DictDataResponse `json:"list"`
	Total int64               `json:"total"`
}
