package service

import (
	"context"
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/model"
	"gin-casbin-admin/internal/repository"
	"github.com/casbin/casbin/v2"
	"slices"
	"strconv"
)

type RoleService interface {
	// Add 添加角色
	Add(ctx context.Context, req *v1.RoleAddRequest) error
	// Delete 删除角色
	Delete(ctx context.Context, req *v1.RoleDeleteRequest) error
	// Update 更新角色
	Update(ctx context.Context, req *v1.RoleUpdateRequest) error
	// Get 获取角色
	Get(ctx context.Context, req *v1.RoleGetRequest) (*v1.RoleGetResponse, error)
	// List 获取角色列表
	List(ctx context.Context, req *v1.RoleListRequest) (*v1.RoleListResponse, error)
}

func NewRoleService(
	service *Service,
	roleRepo repository.RoleRepository,
	menuRepo repository.MenuRepository,
	enforcer *casbin.Enforcer,
) RoleService {
	return &roleService{
		Service:  service,
		roleRepo: roleRepo,
		enforcer: enforcer,
		menuRepo: menuRepo,
	}
}

type roleService struct {
	*Service
	roleRepo repository.RoleRepository
	menuRepo repository.MenuRepository
	enforcer *casbin.Enforcer
}

func (r *roleService) Add(ctx context.Context, req *v1.RoleAddRequest) error {
	err := r.tm.Transaction(ctx, func(ctx context.Context) error {
		var (
			id  int
			err error
		)
		if id, err = r.roleRepo.Create(ctx, &model.AdminRole{
			Name:        req.RoleName,
			Tag:         req.RoleTag,
			Status:      req.Status,
			Description: req.Description,
		}); err != nil {
			return err
		}
		var policies [][]string
		if slices.Contains(req.MenuIds, "*") {
			policies = append(policies, []string{strconv.Itoa(id), "*", "*", "*"})
		} else {
			menus, err := r.menuRepo.GetByIds(ctx, req.MenuIds)
			if err != nil {
				return err
			}
			for _, menu := range menus {
				policies = append(policies, []string{strconv.Itoa(id), strconv.Itoa(int(menu.Id)), strconv.Itoa(int(menu.Type))})
			}
		}
		if _, err := r.enforcer.AddPoliciesEx(policies); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *roleService) Delete(ctx context.Context, req *v1.RoleDeleteRequest) error {
	err := r.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := r.roleRepo.Delete(ctx, req.Id); err != nil {
			return err
		}
		if _, err := r.enforcer.RemoveFilteredPolicy(0, strconv.Itoa(req.Id)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *roleService) Update(ctx context.Context, req *v1.RoleUpdateRequest) error {
	err := r.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := r.roleRepo.Update(ctx, &model.AdminRole{
			Id:          req.Id,
			Name:        req.RoleName,
			Status:      req.Status,
			Description: req.Description,
		}); err != nil {
			return err
		}
		if _, err := r.enforcer.RemoveFilteredPolicy(0, strconv.Itoa(req.Id)); err != nil {
			return err
		}
		menus, err := r.menuRepo.GetByIds(ctx, req.MenuIds)
		if err != nil {
			return err
		}
		policies := make([][]string, 0, len(menus))
		for _, menu := range menus {
			policies = append(policies, []string{strconv.Itoa(req.Id), strconv.Itoa(int(menu.Id)), strconv.Itoa(int(menu.Type))})
		}
		if _, err := r.enforcer.AddPoliciesEx(policies); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *roleService) Get(ctx context.Context, req *v1.RoleGetRequest) (*v1.RoleGetResponse, error) {
	role, err := r.roleRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	policies, err := r.enforcer.GetPermissionsForUser(strconv.Itoa(role.Id))
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(policies))
	for _, p := range policies {
		ids = append(ids, p[3])
	}
	menus, err := r.menuRepo.GetByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	rolePermissions := make([]*v1.RoleMenu, 0, len(policies))
	for _, p := range menus {
		rolePermissions = append(rolePermissions, &v1.RoleMenu{
			Id:     int(p.Id),
			Path:   p.Path,
			Method: p.Method,
			Title:  p.Title,
		})
	}
	return &v1.RoleGetResponse{
		Id:          role.Id,
		RoleName:    role.Name,
		Status:      role.Status,
		Description: role.Description,
		Menus:       rolePermissions,
	}, nil
}

func (r *roleService) List(ctx context.Context, req *v1.RoleListRequest) (*v1.RoleListResponse, error) {
	roles, total, err := r.roleRepo.Pagination(ctx, req.PageNum, req.PageSize, nil)
	if err != nil {
		return nil, err
	}
	res := make([]*v1.RoleGetResponse, 0, len(roles))
	for _, role := range roles {
		res = append(res, &v1.RoleGetResponse{
			Id:          role.Id,
			RoleName:    role.Name,
			Status:      role.Status,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
		})
	}
	return &v1.RoleListResponse{
		Total: total,
		List:  res,
	}, nil
}
