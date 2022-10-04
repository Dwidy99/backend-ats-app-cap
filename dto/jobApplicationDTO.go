package dto

type CreateJobApplicationDTO struct {
	JobID uint64 `json:"job_id" form:"job_id" binding:"required"`
}

type GetJobApplicationEmployee struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateJobApplicationEmployeeDTO struct {
	Status string `json:"status" form:"status" binding:"required"`
}
