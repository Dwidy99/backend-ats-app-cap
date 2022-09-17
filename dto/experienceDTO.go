package dto

type CreateExperienceDTO struct {
	CompanyName string `json:"company_name" form:"company_name" binding:"required"`
	Role        string `json:"role" form:"role" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	DateStart   string `json:"date_start" form:"date_start" binding:"required"`
	DateEnd     string `json:"date_end" form:"date_end" binding:"required"`
	Status      string `json:"status" form:"status" binding:"required"`
}

type GetExperienceDetailDTO struct {
	ID int `uri:"id" binding:"required"`
}