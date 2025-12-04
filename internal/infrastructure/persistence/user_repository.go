package persistence

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/role"
	"github.com/wxlbd/gin-casbin-admin/internal/domain/user"
	"github.com/wxlbd/gin-casbin-admin/internal/infrastructure/persistence/models"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) toEntity(m *models.User) *user.User {
	if m == nil {
		return nil
	}
	return &user.User{
		ID:             m.ID,
		Username:       m.Username,
		Password:       m.Password,
		Nickname:       m.Nickname,
		Phone:          m.Phone,
		Email:          m.Email,
		Avatar:         m.Avatar,
		Status:         m.Status,
		UserType:       m.UserType,
		Signed:         m.Signed,
		LoginIp:        m.LoginIp,
		LoginTime:      m.LoginTime,
		BackendSetting: m.BackendSetting,
		CreatedBy:      m.CreatedBy,
		UpdatedBy:      m.UpdatedBy,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		Remark:         m.Remark,
	}
}

func (r *userRepository) toModel(e *user.User) *models.User {
	if e == nil {
		return nil
	}
	return &models.User{
		ID:             e.ID,
		Username:       e.Username,
		Password:       e.Password,
		Nickname:       e.Nickname,
		Phone:          e.Phone,
		Email:          e.Email,
		Avatar:         e.Avatar,
		Status:         e.Status,
		UserType:       e.UserType,
		Signed:         e.Signed,
		LoginIp:        e.LoginIp,
		LoginTime:      e.LoginTime,
		BackendSetting: e.BackendSetting,
		CreatedBy:      e.CreatedBy,
		UpdatedBy:      e.UpdatedBy,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		Remark:         e.Remark,
	}
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	m := r.toModel(u)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	u.ID = m.ID
	return nil
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	m := r.toModel(u)
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *userRepository) Delete(ctx context.Context, ids ...uint64) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, ids).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint64) (*user.User, error) {
	var m models.User
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toEntity(&m), nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	var m models.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toEntity(&m), nil
}

func (r *userRepository) List(ctx context.Context, query *user.UserQuery) ([]*user.User, int64, error) {
	var modelsList []*models.User
	var total int64
	db := r.db.WithContext(ctx).Model(&models.User{})

	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+query.Phone+"%")
	}
	if query.Email != "" {
		db = db.Where("email LIKE ?", "%"+query.Email+"%")
	}
	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}
	if query.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+query.Nickname+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Offset(offset).Limit(query.PageSize).Order(query.OrderBy).Find(&modelsList).Error; err != nil {
		return nil, 0, err
	}

	var entities []*user.User
	for _, m := range modelsList {
		entities = append(entities, r.toEntity(m))
	}
	return entities, total, nil
}

func (r *userRepository) GetUserRoles(ctx context.Context, userID uint64) ([]*role.Role, error) {
	var roleModels []*models.Role
	err := r.db.WithContext(ctx).
		Table("roles").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roleModels).Error
	if err != nil {
		return nil, err
	}

	var roles []*role.Role
	for _, m := range roleModels {
		roles = append(roles, &role.Role{
			ID:        m.ID,
			Name:      m.Name,
			Code:      m.Code,
			Status:    m.Status,
			Sort:      m.Sort,
			Remark:    m.Remark,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}
	return roles, nil
}

func (r *userRepository) AssignRoles(ctx context.Context, userID uint64, roleIDs []uint64) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete existing roles
		if err := tx.Where("user_id = ?", userID).Delete(&models.UserRole{}).Error; err != nil {
			return err
		}

		// Insert new roles
		if len(roleIDs) == 0 {
			return nil
		}

		var userRoles []models.UserRole
		for _, roleID := range roleIDs {
			userRoles = append(userRoles, models.UserRole{
				UserID: userID,
				RoleID: roleID,
			})
		}

		return tx.Create(&userRoles).Error
	})
}
