package dto

type RegisterApplicantDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterEmployeeDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
