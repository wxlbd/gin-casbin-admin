package menu

import (
	"context"
	"time"
)

// Menu Domain Entity
type Menu struct {
	ID              int64
	ParentID        int64
	MenuType        int32
	Title           string
	Name            string
	Path            string
	Component       string
	Rank            int32
	Redirect        string
	Icon            string
	ExtraIcon       string
	EnterTransition string
	LeaveTransition string
	ActivePath      string
	Auths           string
	FrameSrc        string
	FrameLoading    bool
	KeepAlive       bool
	HiddenTag       bool
	FixedTag        bool
	ShowLink        bool
	ShowParent      bool
	Status          int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Repository Interface
type Repository interface {
	Create(ctx context.Context, menu *Menu) error
	Update(ctx context.Context, menu *Menu) error
	Delete(ctx context.Context, ids ...int64) error
	FindByID(ctx context.Context, id int64) (*Menu, error)
	FindByIDs(ctx context.Context, ids []int64) ([]*Menu, error)
	List(ctx context.Context) ([]*Menu, error)
	FindAll(ctx context.Context) ([]*Menu, error)
	FindByRoleID(ctx context.Context, roleID uint64) ([]*Menu, error)
	FindByUserID(ctx context.Context, userID uint64) ([]*Menu, error)
}
