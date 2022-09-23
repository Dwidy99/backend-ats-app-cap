package dto

type JobDetailDTO struct {
	ID uint64 `uri:"id" binding:"required"`
}

type CreateJobsDTO struct {
	CompanyName       string  `json:"company_name" form:"company_name" binding:"required"`
	Title             string  `json:"title" form:"title" binding:"required"`
	Description       string  `json:"description" form:"description" binding:"required"`
	Location          string  `json:"location" form:"location" binding:"required"`
	Salary            float64 `json:"salary" form:"salary" binding:"required"`
	Type              string  `json:"type" form:"type" binding:"required"`
	LevelOfExperience string  `json:"level_of_experience" form:"level_of_experience" binding:"required"`
	DateStart         string  `json:"date_start" form:"date_start" binding:"required"`
	DateEnd           string  `json:"date_end" form:"date_end" binding:"required"`
}
