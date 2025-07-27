package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255);not null; index:idx_uuid,unique"`
	Name     string    `gorm:"column:name; type:varchar(255);not null"`
	Email    string    `gorm:"column:email; type:varchar(255);not null;unique"`
	Password string    `gorm:"column:password; type:varchar(255);not null"`
	Avatar   string    `gorm:"column:avatar; type:varchar(255)"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
	IsActive bool      `gorm:"column:is_active; type:boolean"`
}

func (u *User) TableName() string {
	return "users"

}
