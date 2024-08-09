package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// AdminRole 后台角色
type AdminRole struct {
	Id          int    `json:"id"`
	Name        string `json:"name" gorm:"type:VARCHAR(20);not null;comment:'角色名称''"`
	Tag         string `json:"tag" gorm:"type:VARCHAR(20);NOT NULL;DEFAULT:'';uniqueIndex:uidx_role_name;"`
	Status      int    `json:"status" gorm:"type:TINYINT;NOT NULL;DEFAULT:1"`
	Description string `json:"description" gorm:"type:VARCHAR(255);NOT NULL;DEFAULT:''"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   soft_delete.DeletedAt `gorm:"uniqueIndex:uidx_role_name" json:"deleted_at"`
}

func (*AdminRole) TableName() string {
	return "admin_roles"
}
