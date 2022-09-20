package entity

import "time"

type Company struct {
	ID        uint64    `json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Address   string    `json:"address"`
	Contact   string    `gorm:"type:varchar(255)" json:"contact"`
	Website   string    `gorm:"type:varchar(255)" json:"website"`
	CreatedAt time.Time `json:"create_at"`
}
