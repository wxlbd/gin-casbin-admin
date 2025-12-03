package user

import (
	"context"
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
)

// User Domain Entity
type User struct {
	ID             uint64
	Username       string
	Password       string
	Nickname       string
	Phone          string
	Email          string
	Avatar         string
	Status         int8
	UserType       int
	Signed         string
	LoginIp        string
	LoginTime      time.Time
	BackendSetting *types.BackendSetting
	CreatedBy      uint64
	UpdatedBy      uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Remark         string
}

// Repository Interface
type Repository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	List(ctx context.Context, query *UserQuery) ([]*User, int64, error)
	GetUserRoles(ctx context.Context, userID uint64) ([]*role.Role, error)
}

type UserQuery struct {
	Username string
	Phone    string
	Email    string
	Status   int8
	Nickname string
	Page     int
	PageSize int
	OrderBy  string
}
