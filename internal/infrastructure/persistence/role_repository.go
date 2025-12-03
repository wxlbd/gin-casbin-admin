package persistence

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/infrastructure/persistence/models"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) role.Repository {
	return &roleRepository{db: db}
}

func (r *roleRepository) toEntity(m *models.Role) *role.Role {
	if m == nil {
		return nil
	}
	return &role.Role{
		ID:        m.ID,
		Name:      m.Name,
		Code:      m.Code,
		Status:    m.Status,
		Sort:      m.Sort,
		Remark:    m.Remark,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r *roleRepository) toModel(e *role.Role) *models.Role {
	if e == nil {
		return nil
	}
	return &models.Role{
		ID:        e.ID,
		Name:      e.Name,
		Code:      e.Code,
		Status:    e.Status,
		Sort:      e.Sort,
		Remark:    e.Remark,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (r *roleRepository) Create(ctx context.Context, e *role.Role) error {
	m := r.toModel(e)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	e.ID = m.ID
	return nil
}

func (r *roleRepository) Update(ctx context.Context, e *role.Role) error {
	m := r.toModel(e)
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *roleRepository) Delete(ctx context.Context, ids ...uint64) error {
	return r.db.WithContext(ctx).Delete(&models.Role{}, ids).Error
}

func (r *roleRepository) FindByID(ctx context.Context, id uint64) (*role.Role, error) {
	var m models.Role
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toEntity(&m), nil
}

func (r *roleRepository) FindByCode(ctx context.Context, code string) (*role.Role, error) {
	var m models.Role
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toEntity(&m), nil
}

func (r *roleRepository) List(ctx context.Context, query *role.RoleQuery) ([]*role.Role, int64, error) {
	var modelsList []*models.Role
	var total int64
	db := r.db.WithContext(ctx).Model(&models.Role{})

	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Code != "" {
		db = db.Where("code LIKE ?", "%"+query.Code+"%")
	}
	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Offset(offset).Limit(query.PageSize).Find(&modelsList).Error; err != nil {
		return nil, 0, err
	}

	var entities []*role.Role
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, total, nil
}

func (r *roleRepository) FindAll(ctx context.Context) ([]*role.Role, error) {
	var modelsList []*models.Role
	if err := r.db.WithContext(ctx).Find(&modelsList).Error; err != nil {
		return nil, err
	}
	var entities []*role.Role
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, nil
}
