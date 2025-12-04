package menu

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/menu"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
)

type Service interface {
	Create(ctx context.Context, req *SysMenuRequest) error
	Update(ctx context.Context, req *SysMenuRequest) error
	Delete(ctx context.Context, ids ...int64) error
	FindByID(ctx context.Context, id int64) (*SysMenuResponse, error)
	List(ctx context.Context) ([]*SysMenuResponse, error)
	GetMenuTree(ctx context.Context) ([]*SysMenuResponse, error)
	GetUserMenuTree(ctx context.Context, userID uint64) ([]*SysMenuResponse, error)
}

type menuService struct {
	repo menu.Repository
}

func NewMenuService(repo menu.Repository) Service {
	return &menuService{repo: repo}
}

func (s *menuService) Create(ctx context.Context, req *SysMenuRequest) error {
	return s.repo.Create(ctx, req.ToEntity())
}

func (s *menuService) Update(ctx context.Context, req *SysMenuRequest) error {
	exist, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.WithMsg(errors.NotFound, "菜单不存在")
		}
		return err
	}
	m := req.ToEntity()
	m.CreatedAt = exist.CreatedAt
	return s.repo.Update(ctx, m)
}

func (s *menuService) Delete(ctx context.Context, ids ...int64) error {
	return s.repo.Delete(ctx, ids...)
}

func (s *menuService) FindByID(ctx context.Context, id int64) (*SysMenuResponse, error) {
	m, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToSysMenuResponse(m), nil
}

func (s *menuService) List(ctx context.Context) ([]*SysMenuResponse, error) {
	menus, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return ToSysMenuList(menus), nil
}

func (s *menuService) GetMenuTree(ctx context.Context) ([]*SysMenuResponse, error) {
	menus, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return buildMenuTree(ToSysMenuList(menus)), nil
}

func (s *menuService) GetUserMenuTree(ctx context.Context, userID uint64) ([]*SysMenuResponse, error) {
	menus, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return buildMenuTree(ToSysMenuList(menus)), nil
}

func buildMenuTree(menus []*SysMenuResponse) []*SysMenuResponse {
	var tree []*SysMenuResponse
	menuMap := make(map[int64]*SysMenuResponse)

	// First pass: map all menus and aggregate button permissions to parents
	for _, m := range menus {
		menuMap[m.ID] = m
	}

	// Second pass: build tree and handle buttons
	for _, m := range menus {
		// Skip buttons in the tree structure, but add their auths to parent
		if m.MenuType == menu.TypeButton {
			if parent, ok := menuMap[m.ParentID]; ok {
				if parent.Auths != "" {
					parent.Auths += "," + m.Auths
				} else {
					parent.Auths = m.Auths
				}
			}
			continue
		}

		if m.ParentID == 0 {
			tree = append(tree, m)
		} else {
			if parent, ok := menuMap[m.ParentID]; ok {
				parent.Children = append(parent.Children, m)
			}
		}
	}
	return tree
}
