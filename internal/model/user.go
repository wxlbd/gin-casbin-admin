package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type User struct {
	Id        uint   `gorm:"primarykey"`
	UserId    string `gorm:"unique;not null;type:VARCHAR(64)"`
	Username  string `gorm:"not null;type:VARCHAR(30);uniqueIndex:udx_username"`
	Password  string `gorm:"not null;type:VARCHAR(255)"`
	Avatar    string `gorm:"not null;type:VARCHAR(255);default:''"`
	Email     string `gorm:"not null;type:VARCHAR(40)"`
	Status    int    `gorm:"not null;type:TINYINT;default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_username"`
}

func (u *User) TableName() string {
	return "users"
}
