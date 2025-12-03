package models

// UserRole 用户角色关联
type UserRole struct {
	UserID uint64 `gorm:"primaryKey"`
	RoleID uint64 `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
