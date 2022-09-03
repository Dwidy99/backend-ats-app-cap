package entity

import "time"

type User struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id_user"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Email       string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Username    string `gorm:"type:varchar(255)" json:"username"`
	Password    string `gorm:"->;<-;not null" json:"-"`
	Token string `gorm:"-" json:"token,omitempty"`
	LevelAccess string `gorm:"type:varchar(255)" json:"level_access"`
	CreatedAt   time.Time `json:"created_at"`
}
