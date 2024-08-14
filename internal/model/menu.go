package model

import (
	"time"
)

type Menu struct {
	ID            uint64 `gorm:"column:id;type:BIGINT UNSIGNED;PRIMARY_KEY;AUTO_INCREMENT;"`
	Title         string `gorm:"column:title;type:VARCHAR(255);not null;default:''"`
	ParentId      uint64 `gorm:"column:parent_id;type:BIGINT UNSIGNED;NOT NULL;DEFAULT:0;comment:父级ID"`
	Type          int8   `gorm:"column:type;type:TINYINT;NOT NULL;DEFAULT:1;comment:类型（1：目录；2：菜单；3：按钮(api)）"`
	Path          string `gorm:"column:path;type:VARCHAR(255);not null;default:'';comment:路由地址"`
	Method        string `gorm:"column:method;type:VARCHAR(255);not null;default:'';comment:请求方式,只有按钮时才生效"`
	ComponentName string `gorm:"column:component_name;type:VARCHAR(255);not null;default:'';comment:组件名称"`
	ComponentPath string `gorm:"column:component_path;type:VARCHAR(255);not null;default:'';comment:组件路径"`
	Redirect      string `gorm:"column:redirect;type:VARCHAR(255);not null;default:'';comment:重定向地址"`
	Icon          string `gorm:"column:icon;type:VARCHAR(255);not null;default:'';comment:菜单图标"`
	IsExternal    int8   `gorm:"column:is_external;type:TINYINT;NOT NULL;DEFAULT:0;comment:是否外链"`
	IsHidden      int8   `gorm:"column:is_hidden;type:TINYINT;NOT NULL;DEFAULT:0;comment:是否隐藏"`
	Sort          int    `gorm:"column:sort;type:TINYINT;NOT NULL;DEFAULT:0;comment:排序"`
	Status        int    `gorm:"column:status;type:TINYINT;NOT NULL;DEFAULT:1;comment:状态（1：启用，2：禁用）"`
	CreateUser    uint64 `gorm:"column:create_user;type:BIGINT UNSIGNED;NOT NULL;DEFAULT:0;comment:创建人"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (*Menu) TableName() string {
	return "sys_menu"
}
