package dict

import (
	"context"
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// DictType Domain Entity
type DictType struct {
	ID        int64
	Code      string
	Name      string
	Status    int32
	Sort      int32
	Remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// DictData Domain Entity
type DictData struct {
	ID        int64
	TypeCode  string
	Label     string
	Value     string
	Status    int32
	Sort      int32
	Remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Repository Interface
type Repository interface {
	// DictType
	CreateType(ctx context.Context, d *DictType) error
	UpdateType(ctx context.Context, d *DictType) error
	DeleteType(ctx context.Context, ids ...int64) error
	FindTypeByID(ctx context.Context, id int64) (*DictType, error)
	FindTypeByCode(ctx context.Context, code string) (*DictType, error)
	ListType(ctx context.Context, query *DictTypeQuery) ([]*DictType, int64, error)

	// DictData
	CreateData(ctx context.Context, d *DictData) error
	UpdateData(ctx context.Context, d *DictData) error
	DeleteData(ctx context.Context, ids ...int64) error
	FindDataByID(ctx context.Context, id int64) (*DictData, error)
	ListData(ctx context.Context, query *DictDataQuery) ([]*DictData, int64, error)
}

type DictTypeQuery struct {
	*types.PageParam
	Name   string
	Code   string
	Status int32
}

type DictDataQuery struct {
	*types.PageParam
	TypeCode string
	Label    string
	Status   int32
}
