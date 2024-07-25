package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type AdminPermissions struct {
	Id        uint64                `gorm:"column:id;type:BIGINT UNSIGNED;PRIMARY_KEY;AUTO_INCREMENT;"`
	Name      string                `gorm:"column:name;type:VARCHAR(255);not null;default:'';uniqueIndex:udx_name"`
	Icon      string                `gorm:"column:icon;type:VARCHAR(255);not null;default:''"`
	Path      string                `gorm:"column:path;type:VARCHAR(255);not null;default:''"`
	Url       string                `gorm:"column:url;type:VARCHAR(255);not null;default:''"`
	Title     string                `gorm:"column:title;type:VARCHAR(255);not null;default:''"`
	Hidden    int                   `gorm:"column:hidden;type:TINYINT;"`
	IsMenu    int                   `gorm:"column:is_menu;type:TINYINT;"`
	PId       int                   `gorm:"column:p_id;type:TINYINT;"`
	Method    string                `gorm:"column:method;type:VARCHAR(255);not null;default:''"`
	Status    int                   `gorm:"column:status;type:TINYINT;NOT NULL"`
	CreatedAt time.Time             `gorm:"column:created_at;type:TIMESTAMP;"`
	UpdatedAt time.Time             `gorm:"column:updated_at;type:TIMESTAMP;"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
}

func (AdminPermissions) TableName() string {
	return "admin_permissions"
}
