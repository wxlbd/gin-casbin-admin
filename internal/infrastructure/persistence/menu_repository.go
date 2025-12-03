package persistence

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/menu"
	"github.com/wxlbd/gin-casbin-admin/internal/infrastructure/persistence/models"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"gorm.io/gorm"
)

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) menu.Repository {
	return &menuRepository{db: db}
}

func (r *menuRepository) toEntity(m *models.SysMenu) *menu.Menu {
	if m == nil {
		return nil
	}
	return &menu.Menu{
		ID:              m.ID,
		ParentID:        m.ParentID,
		MenuType:        m.MenuType,
		Title:           m.Title,
		Name:            m.Name,
		Path:            m.Path,
		Component:       m.Component,
		Rank:            m.Rank,
		Redirect:        m.Redirect,
		Icon:            m.Icon,
		ExtraIcon:       m.ExtraIcon,
		EnterTransition: m.EnterTransition,
		LeaveTransition: m.LeaveTransition,
		ActivePath:      m.ActivePath,
		Auths:           m.Auths,
		FrameSrc:        m.FrameSrc,
		FrameLoading:    m.FrameLoading,
		KeepAlive:       m.KeepAlive,
		HiddenTag:       m.HiddenTag,
		FixedTag:        m.FixedTag,
		ShowLink:        m.ShowLink,
		ShowParent:      m.ShowParent,
		Status:          m.Status,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func (r *menuRepository) toModel(e *menu.Menu) *models.SysMenu {
	if e == nil {
		return nil
	}
	return &models.SysMenu{
		ID:              e.ID,
		ParentID:        e.ParentID,
		MenuType:        e.MenuType,
		Title:           e.Title,
		Name:            e.Name,
		Path:            e.Path,
		Component:       e.Component,
		Rank:            e.Rank,
		Redirect:        e.Redirect,
		Icon:            e.Icon,
		ExtraIcon:       e.ExtraIcon,
		EnterTransition: e.EnterTransition,
		LeaveTransition: e.LeaveTransition,
		ActivePath:      e.ActivePath,
		Auths:           e.Auths,
		FrameSrc:        e.FrameSrc,
		FrameLoading:    e.FrameLoading,
		KeepAlive:       e.KeepAlive,
		HiddenTag:       e.HiddenTag,
		FixedTag:        e.FixedTag,
		ShowLink:        e.ShowLink,
		ShowParent:      e.ShowParent,
		Status:          e.Status,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}

func (r *menuRepository) Create(ctx context.Context, e *menu.Menu) error {
	m := r.toModel(e)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	e.ID = m.ID
	return nil
}

func (r *menuRepository) Update(ctx context.Context, e *menu.Menu) error {
	m := r.toModel(e)
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *menuRepository) Delete(ctx context.Context, ids ...int64) error {
	return r.db.WithContext(ctx).Delete(&models.SysMenu{}, ids).Error
}

func (r *menuRepository) FindByID(ctx context.Context, id int64) (*menu.Menu, error) {
	var m models.SysMenu
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toEntity(&m), nil
}

func (r *menuRepository) FindByIDs(ctx context.Context, ids []int64) ([]*menu.Menu, error) {
	var modelsList []*models.SysMenu
	if err := r.db.WithContext(ctx).Find(&modelsList, ids).Error; err != nil {
		return nil, err
	}
	var entities []*menu.Menu
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, nil
}

func (r *menuRepository) List(ctx context.Context) ([]*menu.Menu, error) {
	return r.FindAll(ctx)
}

func (r *menuRepository) FindAll(ctx context.Context) ([]*menu.Menu, error) {
	var modelsList []*models.SysMenu
	if err := r.db.WithContext(ctx).Order("rank").Find(&modelsList).Error; err != nil {
		return nil, err
	}
	var entities []*menu.Menu
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, nil
}

func (r *menuRepository) FindByRoleID(ctx context.Context, roleID uint64) ([]*menu.Menu, error) {
	// Need to join with role_menus table
	// Assuming there is a role_menus table, but I haven't migrated it yet.
	// For now, I will implement a raw query or use a join if I had the model.
	// Since I haven't migrated RoleMenu model, I'll use a raw query or defer this.
	// Let's check the original implementation.
	// Original used `r.RoleMenu().FindMenusByRoleID`.
	// I should probably migrate RoleMenu model too or handle it here.

	var modelsList []*models.SysMenu
	err := r.db.WithContext(ctx).
		Joins("JOIN role_menus ON role_menus.menu_id = sys_menus.id").
		Where("role_menus.role_id = ?", roleID).
		Find(&modelsList).Error

	if err != nil {
		return nil, err
	}

	var entities []*menu.Menu
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, nil
}

func (r *menuRepository) FindByUserID(ctx context.Context, userID uint64) ([]*menu.Menu, error) {
	var modelsList []*models.SysMenu
	err := r.db.WithContext(ctx).
		Distinct("sys_menus.*").
		Joins("JOIN role_menus ON role_menus.menu_id = sys_menus.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_menus.role_id").
		Where("user_roles.user_id = ?", userID).
		Order("sys_menus.rank").
		Find(&modelsList).Error

	if err != nil {
		return nil, err
	}

	var entities []*menu.Menu
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, nil
}
