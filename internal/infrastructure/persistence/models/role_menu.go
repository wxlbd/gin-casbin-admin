package models

import "time"

// RoleMenu 角色菜单关联表
type RoleMenu struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleID    uint64    `gorm:"column:role_id;not null;index" json:"role_id"`
	MenuID    uint64    `gorm:"column:menu_id;not null;index" json:"menu_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 指定表名
func (RoleMenu) TableName() string {
	return "role_menus"
}
