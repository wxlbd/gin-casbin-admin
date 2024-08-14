package repository

import (
	"context"
	"errors"
	"fmt"
	"gin-casbin-admin/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *model.Menu) error
	Delete(ctx context.Context, ids []uint64) error
	Update(ctx context.Context, menu *model.Menu) error
	Get(ctx context.Context, id uint64) (*model.Menu, error)
	List(ctx context.Context, pageNumber, pageSize int, filters map[string]any) (int64, []*model.Menu, error)
	GetAll(ctx context.Context) ([]*model.Menu, error)
}

type menuRepository struct {
	*Repository
}

func (m *menuRepository) GetAll(ctx context.Context) ([]*model.Menu, error) {
	var list []*model.Menu
	err := m.db.WithContext(ctx).Find(&list).Error
	return list, err
}

func (m *menuRepository) Create(ctx context.Context, menu *model.Menu) error {
	return m.db.WithContext(ctx).Create(&menu).Error
}

func (m *menuRepository) Delete(ctx context.Context, ids []uint64) error {
	return m.db.WithContext(ctx).Where("id in ?", ids).Delete(&model.Menu{}).Error
}

func (m *menuRepository) Update(ctx context.Context, menu *model.Menu) error {
	return m.db.WithContext(ctx).Model(&model.Menu{}).Where("id = ?", menu.ID).Updates(menu).Error
}

func (m *menuRepository) Get(ctx context.Context, id uint64) (*model.Menu, error) {
	var r model.Menu
	if err := m.db.WithContext(ctx).Where("id = ?", id).First(&r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &r, nil
}

func (m *menuRepository) List(ctx context.Context, pageNumber, pageSize int, filters map[string]any) (total int64, list []*model.Menu, err error) {
	db := m.db.WithContext(ctx).Model(&model.Menu{})
	if len(filters) > 0 {
		for k, v := range filters {
			db = db.Where(fmt.Sprintf("%s = ?", k), v)
		}
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if err = db.Offset((pageNumber - 1) * pageSize).Limit(pageSize).Order(
		clause.OrderBy{Columns: []clause.OrderByColumn{
			{Column: clause.Column{Name: "sort"}, Desc: true},
			{Column: clause.Column{Name: "created_at"}, Desc: true},
		}}).Order("created_at desc").Find(&list).Error; err != nil {
		return
	}
	return
}

func NewMenuRepository(repository *Repository) MenuRepository {
	return &menuRepository{
		Repository: repository,
	}
}
