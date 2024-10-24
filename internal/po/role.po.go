package po

import (
	"gorm.io/gorm"
)

// Role represents the role table in the database
type Role struct {
	gorm.Model
	ID    int64  `gorm:"type:int;primary_key;autoIncrement;column:id"`
	Name  string `gorm:"type:varchar(100);column:name"` // Role name
	Note  string `gorm:"type:text;column:note"`         // Role note
	Users []User `gorm:"many2many:go_user_roles;"`      // Many-to-many relationship with User
}

// TableName overrides the default table name to "go_db_role"
func (r *Role) TableName() string {
	return "go_db_role"
}
