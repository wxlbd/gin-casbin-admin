package models

import (
	"time"
)

// Role 角色模型
type Role struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:64"`
	Code      string    `json:"code" gorm:"uniqueIndex;size:64"`
	Status    int8      `json:"status" gorm:"default:1"` // 1: 正常, 2: 禁用
	Sort      int16     `json:"sort" gorm:"default:0"`   // 排序，值越小越靠前
	Remark    string    `json:"remark" gorm:"size:255"`  // 备注
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "roles"
}
