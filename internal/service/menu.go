package service

import (
	"context"
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/model"
	"gin-casbin-admin/internal/repository"
	"go.uber.org/zap"
)

type MenuService interface {
	// Add 添加菜单
	Add(ctx context.Context, req *v1.MenuAddRequest) error
	// Delete 删除菜单
	Delete(ctx context.Context, ids []uint64) error
	// Update 更新菜单
	Update(ctx context.Context, req *v1.MenuUpdateRequest) error
	// Get 获取菜单
	Get(ctx context.Context, req *v1.MenuGetRequest) (*v1.MenuGetResponse, error)
	// List 获取菜单列表
	List(ctx context.Context, req *v1.MenuListRequest) (*v1.MenuListResponse, error)
	// Tree 获取菜单树
	Tree(ctx context.Context, req *v1.MenuTreeRequest) (*v1.MenuTreeResponse, error)
}

type menuService struct {
	Service  *Service
	menuRepo repository.MenuRepository
}

func (m *menuService) Tree(ctx context.Context, req *v1.MenuTreeRequest) (*v1.MenuTreeResponse, error) {
	menus, err := m.menuRepo.GetAll(ctx)
	if err != nil {
		m.Service.logger.Error("获取菜单列表失败", zap.Error(err))
		return &v1.MenuTreeResponse{}, err
	}
	return &v1.MenuTreeResponse{
		List: buildMenuTree(menus, 0),
	}, nil
}

func buildMenuTree(menus []*model.Menu, parentId uint64) []*v1.Menu {
	var tree []*v1.Menu

	for _, v := range menus {
		if v.ParentId == parentId {
			node := &v1.Menu{
				Id:            v.Id,
				Type:          v.Type,
				ParentId:      v.ParentId,
				Icon:          v.Icon,
				Title:         v.Title,
				Path:          v.Path,
				Method:        v.Method,
				Redirect:      v.Redirect,
				ComponentName: v.ComponentName,
				ComponentPath: v.ComponentPath,
				IsHidden:      v.IsHidden,
				IsExternal:    v.IsExternal,
				Sort:          v.Sort,
				Status:        v.Status,
			}
			node.Children = buildMenuTree(menus, v.Id)
			tree = append(tree, node)
		}
	}
	return tree
}

func (m *menuService) Add(ctx context.Context, req *v1.MenuAddRequest) error {
	if err := m.menuRepo.Create(ctx, &model.Menu{
		Title:         req.Title,
		ParentId:      req.ParentId,
		Type:          req.Type,
		Path:          req.Path,
		Method:        req.Method,
		ComponentName: req.ComponentName,
		ComponentPath: req.ComponentPath,
		Redirect:      req.Redirect,
		Icon:          req.Redirect,
		IsExternal:    req.IsExternal,
		IsHidden:      req.IsHidden,
		Sort:          req.Sort,
		Status:        req.Status,
	}); err != nil {
		m.Service.logger.WithContext(ctx).Error("添加菜单失败", zap.Error(err))
		return err
	}
	return nil
}

func (m *menuService) Delete(ctx context.Context, ids []uint64) error {
	err := m.menuRepo.Delete(ctx, ids)
	if err != nil {
		m.Service.logger.WithContext(ctx).Error("删除菜单失败", zap.Error(err))
		return err
	}
	return nil
}

func (m *menuService) Update(ctx context.Context, req *v1.MenuUpdateRequest) error {
	err := m.menuRepo.Update(ctx, &model.Menu{
		Id:            req.Id,
		Title:         req.Title,
		ParentId:      req.ParentId,
		Type:          req.Type,
		Path:          req.Path,
		Method:        req.Method,
		ComponentName: req.ComponentName,
		ComponentPath: req.ComponentPath,
		Redirect:      req.Redirect,
		Icon:          req.Redirect,
		IsExternal:    req.IsExternal,
		IsHidden:      req.IsHidden,
		Sort:          req.Sort,
		Status:        req.Status,
	})
	if err != nil {
		m.Service.logger.WithContext(ctx).Error("更新菜单失败", zap.Error(err))
		return err
	}
	return nil
}

func (m *menuService) Get(ctx context.Context, req *v1.MenuGetRequest) (*v1.MenuGetResponse, error) {
	menu, err := m.menuRepo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.MenuGetResponse{
		Menu: &v1.Menu{
			Id:            menu.Id,
			Title:         menu.Title,
			ParentId:      menu.ParentId,
			Type:          menu.Type,
			Path:          menu.Path,
			Method:        menu.Method,
			ComponentName: menu.ComponentName,
			ComponentPath: menu.ComponentPath,
			Redirect:      menu.Redirect,
			Icon:          menu.Icon,
			IsExternal:    menu.IsExternal,
			IsHidden:      menu.IsHidden,
			Sort:          menu.Sort,
			Status:        menu.Status,
		},
	}, nil
}

func (m *menuService) List(ctx context.Context, req *v1.MenuListRequest) (*v1.MenuListResponse, error) {
	filters := make(map[string]any)
	if req.Title != "" {
		filters["title"] = req.Title
	}
	if req.Type != 0 {
		filters["type"] = req.Type
	}
	total, list, err := m.menuRepo.List(ctx, req.PageNum, req.PageSize, filters)
	if err != nil {
		m.Service.logger.WithContext(ctx).Error("获取菜单列表失败", zap.Error(err))
		return nil, err
	}
	l := make([]*v1.Menu, 0, len(list))
	for _, v := range list {
		l = append(l, &v1.Menu{
			Id:            v.Id,
			Type:          v.Type,
			ParentId:      v.ParentId,
			Icon:          v.Icon,
			Title:         v.Title,
			Path:          v.Path,
			Method:        v.Method,
			Redirect:      v.Redirect,
			ComponentName: v.ComponentName,
			ComponentPath: v.ComponentPath,
			IsHidden:      v.IsHidden,
			IsExternal:    v.IsExternal,
			Sort:          v.Sort,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}
	return &v1.MenuListResponse{
		Total: total,
		List:  l,
	}, nil
}

func NewMenuService(
	service *Service,
	menuRepo repository.MenuRepository,
) MenuService {
	return &menuService{
		Service:  service,
		menuRepo: menuRepo,
	}
}
