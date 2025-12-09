package role

import (
	"context"
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/menu"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/types"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
)

type Service interface {
	Create(ctx context.Context, req *RoleRequest) error
	Update(ctx context.Context, req *RoleRequest) error
	Delete(ctx context.Context, ids ...uint64) error
	FindByID(ctx context.Context, id uint64) (*RoleResponse, error)
	List(ctx context.Context, req *RoleListRequest) (*RoleListResponse, error)
	GetAllRoles(ctx context.Context) ([]*RoleResponse, error)
	AssignMenuByIds(ctx context.Context, roleID uint64, menuIds []uint64) error
	GetPermittedMenus(ctx context.Context, roleID uint64) ([]*menu.Menu, error)
}

type roleService struct {
	repo     role.Repository
	menuRepo menu.Repository
	enforcer *casbin.Enforcer
}

func NewRoleService(repo role.Repository, menuRepo menu.Repository, enforcer *casbin.Enforcer) Service {
	return &roleService{
		repo:     repo,
		menuRepo: menuRepo,
		enforcer: enforcer,
	}
}

func (s *roleService) Create(ctx context.Context, req *RoleRequest) error {
	if s.IsCodeExists(ctx, req.Code) {
		return errors.WithMsg(errors.AlreadyExists, "角色代码已存在")
	}
	return s.repo.Create(ctx, req.ToEntity())
}

func (s *roleService) IsCodeExists(ctx context.Context, code string) bool {
	r, _ := s.repo.FindByCode(ctx, code)
	return r != nil
}

func (s *roleService) Update(ctx context.Context, req *RoleRequest) error {
	existRole, err := s.repo.FindByID(ctx, req.ID)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.WithMsg(errors.NotFound, "角色不存在")
		}
		return err
	}

	if req.Code != existRole.Code {
		if s.IsCodeExists(ctx, req.Code) {
			return errors.WithMsg(errors.AlreadyExists, "角色代码已存在")
		}
	}

	r := req.ToEntity()
	// 保持原有的创建时间
	r.CreatedAt = existRole.CreatedAt
	return s.repo.Update(ctx, r)
}

func (s *roleService) Delete(ctx context.Context, ids ...uint64) error {
	for _, id := range ids {
		role, err := s.repo.FindByID(ctx, id)
		if err != nil {
			if err == errors.ErrNotFound {
				continue
			}
			return err
		}
		// 删除权限
		_, err = s.enforcer.DeletePermissionsForUser(role.Code)
		if err != nil {
			return err
		}
	}
	return s.repo.Delete(ctx, ids...)
}

func (s *roleService) FindByID(ctx context.Context, id uint64) (*RoleResponse, error) {
	r, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToRoleResponse(r), nil
}

func (s *roleService) List(ctx context.Context, req *RoleListRequest) (*RoleListResponse, error) {
	query := req.ToQuery()
	roles, total, err := s.repo.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return &RoleListResponse{
		List:  ToRoleList(roles),
		Total: total,
	}, nil
}

func (s *roleService) GetAllRoles(ctx context.Context) ([]*RoleResponse, error) {
	roles, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return ToRoleList(roles), nil
}

func (s *roleService) AssignMenuByIds(ctx context.Context, roleID uint64, menuIds []uint64) error {
	// 1. 获取菜单列表
	var menuIdsInt64 []int64
	for _, id := range menuIds {
		menuIdsInt64 = append(menuIdsInt64, int64(id))
	}

	menusList, err := s.menuRepo.FindByIDs(ctx, menuIdsInt64)
	if err != nil {
		return err
	}

	role, err := s.repo.FindByID(ctx, roleID)
	if err != nil {
		if err == errors.ErrNotFound {
			return errors.WithMsg(errors.NotFound, "角色不存在")
		}
		return err
	}

	// 2. 更新Casbin权限策略
	// 删除旧权限
	if _, err := s.enforcer.DeletePermissionsForUser(role.Code); err != nil {
		return err
	}

	// 添加新权限
	for _, m := range menusList {
		if types.MenuType(m.MenuType) == types.MenuTypeButton {
			path, method := convertMenuToAPI(m.Auths)
			if path != "" && method != "" {
				_, err = s.enforcer.AddPolicy(role.Code, path, method)
				if err != nil {
					return err
				}
			}
		}
	}

	// 3. 更新角色-菜单关联关系
	return s.repo.UpdateMenus(ctx, roleID, menuIds)
}

func (s *roleService) GetPermittedMenus(ctx context.Context, roleID uint64) ([]*menu.Menu, error) {
	role, err := s.repo.FindByID(ctx, roleID)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, errors.WithMsg(errors.NotFound, "角色不存在")
		}
		return nil, err
	}

	if role.Code == "SuperAdmin" {
		return s.menuRepo.FindAll(ctx)
	}

	return s.menuRepo.FindByRoleID(ctx, roleID)
}

// convertMenuToAPI 将菜单权限标识转换为API路径和方法
func convertMenuToAPI(menuName string) (path, method string) {
	const apiPrefix = "/api/v1" // 添加 v1 版本前缀
	parts := strings.Split(menuName, ":")
	if len(parts) < 3 {
		return "", ""
	}

	module := parts[0]   // system
	resource := parts[1] // dict 或 user
	var action string
	var subResource string

	// 处理四段式权限标识: system:dict:type:list
	if len(parts) >= 4 {
		subResource = parts[2] // type
		action = parts[3]      // list
	} else {
		action = parts[2] // list
	}

	actionMap := map[string]struct {
		method     string
		pathSuffix string
	}{
		"create":  {"POST", ""},
		"save":    {"POST", ""},
		"update":  {"PUT", ":id"},
		"delete":  {"DELETE", ":ids"},
		"get":     {"GET", ":id"},
		"detail":  {"GET", ":id"},
		"list":    {"GET", ""},
		"index":   {"GET", ""},
		"enable":  {"PATCH", "enable"},
		"disable": {"PATCH", "disable"},
		"assign":  {"POST", "assign"},
		"revoke":  {"POST", "revoke"},
		"upload":  {"POST", "upload"},
		"export":  {"GET", "export"},
		"import":  {"POST", "import"},
		"batch":   {"POST", "batch"},
		"tree":    {"GET", "tree"},
		"status":  {"PATCH", "status"},
		"set":     {"PUT", ":id"},
	}
	item, ok := actionMap[action]
	if !ok {
		return "", ""
	}

	// 构建路径
	if subResource != "" {
		// 四段式: /api/v1/system/dict-type
		path = fmt.Sprintf("%s/%s/%s-%s", apiPrefix, module, resource, subResource)
	} else {
		// 三段式: /api/v1/system/user
		path = fmt.Sprintf("%s/%s/%s", apiPrefix, module, resource)
	}

	// 添加路径后缀（如 :id, :ids）
	if item.pathSuffix != "" {
		path = fmt.Sprintf("%s/%s", path, item.pathSuffix)
	}

	return path, item.method
}
