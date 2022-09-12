package dto

type RegisterApplicantDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type RegisterEmployeeDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Role     string `json:"role" form:"role" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
