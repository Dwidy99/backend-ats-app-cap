package entity

type Employee struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Contact string `gorm:"type:varchar(255)" json:"contact"`
}
