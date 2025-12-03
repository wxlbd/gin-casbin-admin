package menu

import (
	"time"

	"github.com/wxlbd/gin-casbin-admin/internal/domain/menu"
)

// SysMenuRequest 创建/更新菜单请求
type SysMenuRequest struct {
	ID              int64  `json:"id"`
	ParentID        int64  `json:"parentId"`                 // 父菜单ID
	MenuType        int32  `json:"menuType"`                 // 菜单类型（1代表菜单、2代表iframe、3代表外链、4代表按钮）
	Title           string `json:"title" binding:"required"` // 菜单名称
	Name            string `json:"name"`                     // 路由名称
	Path            string `json:"path"`                     // 路由路径
	Component       string `json:"component"`                // 组件路径
	Rank            int32  `json:"rank"`                     // 显示排序
	Redirect        string `json:"redirect"`                 // 重定向地址
	Icon            string `json:"icon"`                     // 菜单图标
	ExtraIcon       string `json:"extraIcon"`                // 右侧图标
	EnterTransition string `json:"enterTransition"`          // 进场动画
	LeaveTransition string `json:"leaveTransition"`          // 离场动画
	ActivePath      string `json:"activePath"`               // 激活路由路径
	Auths           string `json:"auths"`                    // 权限标识
	FrameSrc        string `json:"frameSrc"`                 // 外链地址
	FrameLoading    bool   `json:"frameLoading"`             // 是否显示加载动画
	KeepAlive       bool   `json:"keepAlive"`                // 是否缓存
	HiddenTag       bool   `json:"hiddenTag"`                // 是否隐藏标签
	FixedTag        bool   `json:"fixedTag"`                 // 是否固定标签
	ShowLink        bool   `json:"showLink"`                 // 是否显示
	ShowParent      bool   `json:"showParent"`               // 是否显示父级菜单
	Status          int32  `json:"status"`                   // 菜单状态（0停用 1正常）
}

func (req *SysMenuRequest) ToEntity() *menu.Menu {
	return &menu.Menu{
		ID:              req.ID,
		ParentID:        req.ParentID,
		MenuType:        req.MenuType,
		Title:           req.Title,
		Name:            req.Name,
		Path:            req.Path,
		Component:       req.Component,
		Rank:            req.Rank,
		Redirect:        req.Redirect,
		Icon:            req.Icon,
		ExtraIcon:       req.ExtraIcon,
		EnterTransition: req.EnterTransition,
		LeaveTransition: req.LeaveTransition,
		ActivePath:      req.ActivePath,
		Auths:           req.Auths,
		FrameSrc:        req.FrameSrc,
		FrameLoading:    req.FrameLoading,
		KeepAlive:       req.KeepAlive,
		HiddenTag:       req.HiddenTag,
		FixedTag:        req.FixedTag,
		ShowLink:        req.ShowLink,
		ShowParent:      req.ShowParent,
		Status:          req.Status,
	}
}

// SysMenuResponse 菜单响应
type SysMenuResponse struct {
	ID              int64              `json:"id"`
	ParentID        int64              `json:"parentId"`
	MenuType        int32              `json:"menuType"`
	Title           string             `json:"title"`
	Name            string             `json:"name"`
	Path            string             `json:"path"`
	Component       string             `json:"component"`
	Rank            int32              `json:"rank"`
	Redirect        string             `json:"redirect"`
	Icon            string             `json:"icon"`
	ExtraIcon       string             `json:"extraIcon"`
	EnterTransition string             `json:"enterTransition"`
	LeaveTransition string             `json:"leaveTransition"`
	ActivePath      string             `json:"activePath"`
	Auths           string             `json:"auths"`
	FrameSrc        string             `json:"frameSrc"`
	FrameLoading    bool               `json:"frameLoading"`
	KeepAlive       bool               `json:"keepAlive"`
	HiddenTag       bool               `json:"hiddenTag"`
	FixedTag        bool               `json:"fixedTag"`
	ShowLink        bool               `json:"showLink"`
	ShowParent      bool               `json:"showParent"`
	Status          int32              `json:"status"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	Children        []*SysMenuResponse `json:"children,omitempty"`
}

func ToSysMenuResponse(m *menu.Menu) *SysMenuResponse {
	if m == nil {
		return nil
	}
	return &SysMenuResponse{
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

func ToSysMenuList(menus []*menu.Menu) []*SysMenuResponse {
	list := make([]*SysMenuResponse, 0, len(menus))
	for _, m := range menus {
		list = append(list, ToSysMenuResponse(m))
	}
	return list
}

// SysMenuListRequest 菜单列表请求
type SysMenuListRequest struct {
	Title    string `form:"title"`    // 菜单名称
	Status   int32  `form:"status"`   // 状态
	MenuType int32  `form:"type"`     // 菜单类型
	Page     int    `form:"page"`     // 页码
	PageSize int    `form:"pageSize"` // 每页数量
}
