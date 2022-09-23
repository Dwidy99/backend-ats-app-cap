package entity

import "time"

type Jobapplication struct {
	ID          uint64 `json:"id"`
	ApplicantID uint64 `json:"applicant_id"`
	JobID       uint64 `json:"job_id"`
	Status      string `json:"status" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at"`
}