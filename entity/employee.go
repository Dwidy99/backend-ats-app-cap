package entity

type Employee struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `gorm:"type:varchar(255)" json:"name"`
	Phone  string `gorm:"type:varchar(255)" json:"phone_number"`
}