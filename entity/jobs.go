package entity

import "time"

type Jobs struct {
	ID                uint64    `json:"id"`
	CompanyID         uint64    `json:"company_id"`
	Title             string    `gorm:"type:varchar(255)" json:"title"`
	Description       string    `json:"description"`
	Location          string    `gorm:"type:varchar(255)" json:"location"`
	Salary            float64   `json:"salary"`
	Type              string    `gorm:"type:varchar(255)" json:"type"`
	LevelOfExperience string    `gorm:"type:varchar(255)" json:"level_of_experience"`
	DateStart         string    `gorm:"type:varchar(255)" json:"date_start"`
	DateEnd           string    `gorm:"type:varchar(255)" json:"date_end"`
	CreatedAt         time.Time `json:"created_at"`
	PostedBy          uint64    `json:"posted_by"`
}
