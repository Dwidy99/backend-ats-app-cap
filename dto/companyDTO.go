package dto

type companyDTO struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Email   string `json:"email" form:"email" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Contact string `json:"contact" form:"contact" binding:"required"`
	Website string `json:"website" form:"website" binding:"required"`
}
