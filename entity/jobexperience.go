package entity

type Jobexperience struct {
	ID          uint64 `json:"id"`
	ApplicantID uint64 `json:"applicant_id"`
	CompanyName string `gorm:"type:varchar(255)" json:"company_name"`
	Role        string `gorm:"type:varchar(255)" json:"role"`
	Description string `json:"description"`
	DateStart   string `gorm:"type:varchar(255)" json:"date_start"`
	DateEnd     string `gorm:"type:varchar(255)" json:"date_end"`
	Status      string `gorm:"type:varchar(255)" json:"status"`
}