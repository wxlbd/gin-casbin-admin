package persistence

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/dict"
	"github.com/wxlbd/gin-casbin-admin/internal/infrastructure/persistence/models"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"gorm.io/gorm"
)

type dictRepository struct {
	db *gorm.DB
}

func NewDictRepository(db *gorm.DB) dict.Repository {
	return &dictRepository{db: db}
}

// --- DictType ---

func (r *dictRepository) toTypeEntity(m *models.DictType) *dict.DictType {
	if m == nil {
		return nil
	}
	return &dict.DictType{
		ID:        m.ID,
		Code:      m.Code,
		Name:      m.Name,
		Status:    m.Status,
		Sort:      m.Sort,
		Remark:    m.Remark,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r *dictRepository) toTypeModel(e *dict.DictType) *models.DictType {
	if e == nil {
		return nil
	}
	return &models.DictType{
		ID:        e.ID,
		Code:      e.Code,
		Name:      e.Name,
		Status:    e.Status,
		Sort:      e.Sort,
		Remark:    e.Remark,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (r *dictRepository) CreateType(ctx context.Context, e *dict.DictType) error {
	m := r.toTypeModel(e)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	e.ID = m.ID
	return nil
}

func (r *dictRepository) UpdateType(ctx context.Context, e *dict.DictType) error {
	m := r.toTypeModel(e)
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *dictRepository) DeleteType(ctx context.Context, ids ...int64) error {
	return r.db.WithContext(ctx).Delete(&models.DictType{}, ids).Error
}

func (r *dictRepository) FindTypeByID(ctx context.Context, id int64) (*dict.DictType, error) {
	var m models.DictType
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toTypeEntity(&m), nil
}

func (r *dictRepository) FindTypeByCode(ctx context.Context, code string) (*dict.DictType, error) {
	var m models.DictType
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toTypeEntity(&m), nil
}

func (r *dictRepository) ListType(ctx context.Context, query *dict.DictTypeQuery) ([]*dict.DictType, int64, error) {
	var modelsList []*models.DictType
	var total int64
	db := r.db.WithContext(ctx).Model(&models.DictType{})

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
	if err := db.Offset(offset).Limit(query.PageSize).Order("sort").Find(&modelsList).Error; err != nil {
		return nil, 0, err
	}

	var entities []*dict.DictType
	for _, m := range modelsList {
		entities = append(entities, r.toTypeEntity(m))
	}
	return entities, total, nil
}

// --- DictData ---

func (r *dictRepository) toDataEntity(m *models.DictDatum) *dict.DictData {
	if m == nil {
		return nil
	}
	return &dict.DictData{
		ID:        m.ID,
		TypeCode:  m.TypeCode,
		Label:     m.Label,
		Value:     m.Value,
		Status:    m.Status,
		Sort:      m.Sort,
		Remark:    m.Remark,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r *dictRepository) toDataModel(e *dict.DictData) *models.DictDatum {
	if e == nil {
		return nil
	}
	return &models.DictDatum{
		ID:        e.ID,
		TypeCode:  e.TypeCode,
		Label:     e.Label,
		Value:     e.Value,
		Status:    e.Status,
		Sort:      e.Sort,
		Remark:    e.Remark,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (r *dictRepository) CreateData(ctx context.Context, e *dict.DictData) error {
	m := r.toDataModel(e)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return err
	}
	e.ID = m.ID
	return nil
}

func (r *dictRepository) UpdateData(ctx context.Context, e *dict.DictData) error {
	m := r.toDataModel(e)
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *dictRepository) DeleteData(ctx context.Context, ids ...int64) error {
	return r.db.WithContext(ctx).Delete(&models.DictDatum{}, ids).Error
}

func (r *dictRepository) FindDataByID(ctx context.Context, id int64) (*dict.DictData, error) {
	var m models.DictDatum
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return r.toDataEntity(&m), nil
}

func (r *dictRepository) ListData(ctx context.Context, query *dict.DictDataQuery) ([]*dict.DictData, int64, error) {
	var modelsList []*models.DictDatum
	var total int64
	db := r.db.WithContext(ctx).Model(&models.DictDatum{})

	if query.TypeCode != "" {
		db = db.Where("type_code = ?", query.TypeCode)
	}
	if query.Label != "" {
		db = db.Where("label LIKE ?", "%"+query.Label+"%")
	}
	if query.Status != 0 {
		db = db.Where("status = ?", query.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	if err := db.Offset(offset).Limit(query.PageSize).Order("sort").Find(&modelsList).Error; err != nil {
		return nil, 0, err
	}

	var entities []*dict.DictData
	for _, m := range modelsList {
		entities = append(entities, r.toDataEntity(m))
	}
	return entities, total, nil
}
