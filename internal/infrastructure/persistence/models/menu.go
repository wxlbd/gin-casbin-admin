package models

import (
	"time"
)

const TableNameSysMenu = "sys_menus"

// SysMenu 菜单权限表
type SysMenu struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:菜单ID" json:"id"`                                 // 菜单ID
	ParentID        int64     `gorm:"column:parent_id;comment:父菜单ID" json:"parent_id"`                                                // 父菜单ID
	MenuType        int32     `gorm:"column:menu_type;not null;default:1;comment:菜单类型（1代表菜单、2代表iframe、3代表外链、4代表按钮）" json:"menu_type"` // 菜单类型（1代表菜单、2代表iframe、3代表外链、4代表按钮）
	Title           string    `gorm:"column:title;not null;comment:菜单名称" json:"title"`                                                // 菜单名称
	Name            string    `gorm:"column:name;comment:路由名称" json:"name"`                                                           // 路由名称
	Path            string    `gorm:"column:path;comment:路由路径" json:"path"`                                                           // 路由路径
	Component       string    `gorm:"column:component;comment:组件路径" json:"component"`                                                 // 组件路径
	Rank            int32     `gorm:"column:rank;default:99;comment:显示排序" json:"rank"`                                                // 显示排序
	Redirect        string    `gorm:"column:redirect;comment:重定向地址" json:"redirect"`                                                  // 重定向地址
	Icon            string    `gorm:"column:icon;comment:菜单图标" json:"icon"`                                                           // 菜单图标
	ExtraIcon       string    `gorm:"column:extra_icon;comment:右侧图标" json:"extra_icon"`                                               // 右侧图标
	EnterTransition string    `gorm:"column:enter_transition;comment:进场动画" json:"enter_transition"`                                   // 进场动画
	LeaveTransition string    `gorm:"column:leave_transition;comment:离场动画" json:"leave_transition"`                                   // 离场动画
	ActivePath      string    `gorm:"column:active_path;comment:激活路由路径" json:"active_path"`                                           // 激活路由路径
	Auths           string    `gorm:"column:auths;comment:权限标识" json:"auths"`                                                         // 权限标识
	FrameSrc        string    `gorm:"column:frame_src;comment:外链地址" json:"frame_src"`                                                 // 外链地址
	FrameLoading    bool      `gorm:"column:frame_loading;default:true;comment:是否显示加载动画" json:"frame_loading"`                        // 是否显示加载动画
	KeepAlive       bool      `gorm:"column:keep_alive;default:false;comment:是否缓存" json:"keep_alive"`                                 // 是否缓存
	HiddenTag       bool      `gorm:"column:hidden_tag;default:false;comment:是否隐藏标签" json:"hidden_tag"`                               // 是否隐藏标签
	FixedTag        bool      `gorm:"column:fixed_tag;default:false;comment:是否固定标签" json:"fixed_tag"`                                 // 是否固定标签
	ShowLink        bool      `gorm:"column:show_link;default:true;comment:是否显示" json:"show_link"`                                    // 是否显示
	ShowParent      bool      `gorm:"column:show_parent;default:false;comment:是否显示父级菜单" json:"show_parent"`                           // 是否显示父级菜单
	Status          int32     `gorm:"column:status;default:1;comment:菜单状态（0停用 1正常）" json:"status"`                                    // 菜单状态（0停用 1正常）
	CreatedAt       time.Time `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                               // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at;comment:更新时间" json:"updated_at"`                                               // 更新时间
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}
