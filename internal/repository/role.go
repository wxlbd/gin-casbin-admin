package repository

import (
	"context"
	"errors"
	"fmt"
	"gin-casbin-admin/internal/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(ctx context.Context, role *model.AdminRole) (int, error)
	Update(ctx context.Context, role *model.AdminRole) error
	GetByID(ctx context.Context, id int) (*model.AdminRole, error)
	FindAll(ctx context.Context, filters map[string]any) ([]*model.AdminRole, error)
	Delete(ctx context.Context, id int) error
	Pagination(ctx context.Context, page, pageSize int, filters map[string]any) ([]*model.AdminRole, int, error)
}

func NewRoleRepository(
	r *Repository,
) RoleRepository {
	return &roleRepository{
		Repository: r,
	}
}

type roleRepository struct {
	*Repository
}

func (r *roleRepository) Delete(ctx context.Context, id int) error {
	return r.DB(ctx).Delete(&model.AdminRole{}, id).Error
}

func (r *roleRepository) Create(ctx context.Context, role *model.AdminRole) (int, error) {
	if err := r.DB(ctx).Create(role).Error; err != nil {
		return 0, err
	}
	return role.Id, nil
}

func (r *roleRepository) Update(ctx context.Context, role *model.AdminRole) error {
	return r.DB(ctx).Save(role).Error
}

func (r *roleRepository) GetByID(ctx context.Context, id int) (*model.AdminRole, error) {
	var role model.AdminRole
	if err := r.DB(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) FindAll(ctx context.Context, filters map[string]any) ([]*model.AdminRole, error) {
	var roles []*model.AdminRole
	query := r.DB(ctx).Model(&model.AdminRole{})
	for k, v := range filters {
		query.Where(fmt.Sprintf("%s = ?", k), v)
	}
	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) Pagination(ctx context.Context, page, pageSize int, filters map[string]any) ([]*model.AdminRole, int, error) {
	var roles []*model.AdminRole
	var count int64
	query := r.DB(ctx).Model(&model.AdminRole{})
	for k, v := range filters {
		query.Where(fmt.Sprintf("%s = ?", k), v)
	}
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	fmt.Println(page, pageSize)
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	return roles, int(count), nil
}
