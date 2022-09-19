package entity

type Jobskill struct {
	ID   uint64 `json:"id"`
	Name string `gorm:"type:varchar(255)" json:"skill_name"`
}

type Jobskillapplicant struct {
	SkillID     uint64 `json:"skill_id"`
	ApplicantID uint64 `json:"applicant_id"`
}