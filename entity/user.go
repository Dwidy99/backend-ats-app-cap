package entity

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"user_id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	Role      string    `gorm:"type:varchar(255)" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
