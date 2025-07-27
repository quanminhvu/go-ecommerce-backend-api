package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment:'primary key with id"`
	RoleName string `gorm:"column:role_name; type:varchar(255);not null; unique"`
	RoleNote string `gorm:"column:role_note; type:text;not null"`
}

func (u *Role) TableName() string {
	return "roles"

}
