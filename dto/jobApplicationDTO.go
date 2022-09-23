package dto

type CreateJobApplication struct {
	JobID  uint64 `json:"job_id" form:"job_id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
}