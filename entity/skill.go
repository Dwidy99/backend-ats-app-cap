package entity

type Jobskill struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}

type Jobskillapplicant struct {
	SkillID     uint64   `json:"skill_id" gorm:"foreignKey:JobskillRefer"`
	ApplicantID uint64   `json:"applicant_id" gorm:"foreignKey:ApplicantRefer"`
	Jobskill    Jobskill `json:"job_skill"`
}

type JobSkillsDetailFormatter struct {
	SkillID     uint64 `json:"skill_id"`
	ApplicantID uint64 `json:"applicant_id"`
	Name        string `json:"name"`
}

type Jobskillrequirement struct {
	SkillID uint64 `json:"id_skill" gorm:"foreignKey:JobskillRefer"`
	JobID   uint64 `json:"id_job" gorm:"foreignKey:JobsRefer"`
}