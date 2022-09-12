package entity

type Applicant struct {
	ID            uint64 `json:"id"`
	UserID        uint64 `json:"user_id"`
	FirstName     string `gorm:"type:varchar(255)" json:"first_name"`
	LastName      string `gorm:"type:varchar(255)" json:"last_name"`
	Avatar        string `gorm:"type:varchar(255)" json:"avatar"`
	Phone         string `gorm:"type:varchar(255)" json:"phone_number"`
	AccountStatus int    `json:"account_status"`
	LastEducation string `gorm:"type:varchar(255)" json:"last_education"`
	LinkedURL     string `gorm:"type:varchar(255)" json:"linkedin_url"`
	GithubURL     string `gorm:"type:varchar(255)" json:"github_url"`
}