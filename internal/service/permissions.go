package service

import (
	"context"
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/model"
	"gin-casbin-admin/internal/repository"
	"go.uber.org/zap"
)

type PermissionsService interface {
	// Add 添加权限
	Add(ctx context.Context, req *v1.PermissionsAddRequest) error
	// Delete 删除权限
	Delete(ctx context.Context, req *v1.PermissionsDeleteRequest) error
	// Update 更新权限
	Update(ctx context.Context, req *v1.PermissionsUpdateRequest) error
	// Get 获取权限
	Get(ctx context.Context, req *v1.PermissionsGetRequest) (*v1.PermissionsGetResponse, error)
	// List 获取权限列表
	List(ctx context.Context, req *v1.PermissionTreeRequest) (*v1.PermissionsTreeResponse, error)
}

func NewPermissionsService(
	service *Service,
	permissionsRepo repository.PermissionsRepository,
) PermissionsService {
	return &permissionsService{
		Service:         service,
		permissionsRepo: permissionsRepo,
	}
}

type permissionsService struct {
	*Service
	permissionsRepo repository.PermissionsRepository
}

func (p *permissionsService) Add(ctx context.Context, req *v1.PermissionsAddRequest) error {
	if err := p.permissionsRepo.Create(ctx, &model.AdminPermissions{
		Name:   req.Name,
		Icon:   req.Icon,
		Path:   req.Path,
		Url:    req.Url,
		Title:  req.Title,
		Hidden: req.Hidden,
		IsMenu: req.IsMenu,
		PId:    req.PId,
		Method: req.Method,
		Status: req.Status,
	}); err != nil {
		p.logger.Error("添加权限失败", zap.Error(err))
		return err
	}
	return nil
}

func (p *permissionsService) Delete(ctx context.Context, req *v1.PermissionsDeleteRequest) error {
	return p.permissionsRepo.Delete(ctx, req.Id)
}

func (p *permissionsService) Update(ctx context.Context, req *v1.PermissionsUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p *permissionsService) Get(ctx context.Context, req *v1.PermissionsGetRequest) (*v1.PermissionsGetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *permissionsService) List(ctx context.Context, req *v1.PermissionTreeRequest) (*v1.PermissionsTreeResponse, error) {
	all, err := p.permissionsRepo.FindAll(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}
	return &v1.PermissionsTreeResponse{
		List: buildTree(all, 0),
	}, nil
}

func buildTree(permissions []*model.AdminPermissions, pid int) []*v1.Permissions {
	var list []*v1.Permissions
	for _, v := range permissions {
		if v.PId == pid {
			p := &v1.Permissions{
				Id:     int(v.Id),
				Name:   v.Name,
				Icon:   v.Icon,
				Path:   v.Path,
				Url:    v.Url,
				Title:  v.Title,
				Hidden: v.Hidden,
				IsMenu: v.IsMenu,
				Pid:    v.PId,
				Method: v.Method,
				Status: v.Status,
			}
			p.Children = buildTree(permissions, int(v.Id))
			list = append(list, p)
		}
	}
	return list
}
