package repository

import (
	"context"
	"fmt"
	"gin-casbin-admin/internal/model"
)

type PermissionsRepository interface {
	Create(ctx context.Context, permissions *model.AdminPermissions) error
	Update(ctx context.Context, permissions *model.AdminPermissions) error
	GetByID(ctx context.Context, id int) (*model.AdminPermissions, error)
	FindAll(ctx context.Context, filters map[string]any) ([]*model.AdminPermissions, error)
	Delete(ctx context.Context, id int) error
	GetByIds(ctx context.Context, ids []string) ([]*model.AdminPermissions, error)
}

func NewPermissionsRepository(
	r *Repository,
) PermissionsRepository {
	return &permissionsRepository{
		Repository: r,
	}
}

type permissionsRepository struct {
	*Repository
}

func (p *permissionsRepository) GetByIds(ctx context.Context, ids []string) ([]*model.AdminPermissions, error) {
	var list []*model.AdminPermissions
	err := p.DB(ctx).Where("id in ?", ids).Find(&list).Error
	return list, err
}

func (p *permissionsRepository) Delete(ctx context.Context, id int) error {
	return p.DB(ctx).Delete(&model.AdminPermissions{}, id).Error
}

func (p *permissionsRepository) Create(ctx context.Context, permissions *model.AdminPermissions) error {
	return p.DB(ctx).Create(permissions).Error
}

func (p *permissionsRepository) Update(ctx context.Context, permissions *model.AdminPermissions) error {
	return p.DB(ctx).Save(permissions).Error
}

func (p *permissionsRepository) GetByID(ctx context.Context, id int) (*model.AdminPermissions, error) {
	var permissions model.AdminPermissions
	if err := p.DB(ctx).First(&permissions, id).Error; err != nil {
		return nil, err
	}
	return &permissions, nil
}

func (p *permissionsRepository) FindAll(ctx context.Context, filters map[string]any) ([]*model.AdminPermissions, error) {
	var permissions []*model.AdminPermissions
	query := p.DB(ctx).Model(&model.AdminPermissions{})
	for k, v := range filters {
		query.Where(fmt.Sprintf("%s = ?", k), v)
	}
	if err := query.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
