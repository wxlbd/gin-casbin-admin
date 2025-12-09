package role

import (
	"context"
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// Role Domain Entity
type Role struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Status    int8      `json:"status"`
	Sort      int16     `json:"sort"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Repository Interface
type Repository interface {
	Create(ctx context.Context, role *Role) error
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*Role, error)
	FindByCode(ctx context.Context, code string) (*Role, error)
	List(ctx context.Context, query *RoleQuery) ([]*Role, int64, error)
	FindAll(ctx context.Context) ([]*Role, error)
	UpdateMenus(ctx context.Context, roleID uint64, menuIDs []uint64) error
}

type RoleQuery struct {
	*types.PageParam
	Name   string
	Code   string
	Status int8
}
