package po

import (
	"gorm.io/gorm"
)

// User represents the user table in the database
type User struct {
	gorm.Model        // Embedded gorm.Model for common fields like ID, CreatedAt, UpdatedAt, DeletedAt
	ID         int64  `gorm:"type:int;primary_key;autoIncrement;column:id"`
	Username   string `gorm:"type:varchar(100);unique;not null;column:username"` // Unique username
	Email      string `gorm:"type:varchar(100);unique;column:email"`             // Unique email
	Phone      string `gorm:"type:varchar(20);column:phone"`                     // Phone number
	Address    string `gorm:"type:text;column:address"`                          // Address
	Active     bool   `gorm:"column:active"`                                     // Active status
	Roles      []Role `gorm:"many2many:go_user_roles;"`                          // Many-to-many relationship with Role
}

// TableName overrides the default table name to "go_db_user"
func (u *User) TableName() string {
	return "go_db_user"
}
